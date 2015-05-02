import urllib
import string
import re
from bs4 import BeautifulSoup

'''URL parameters'''
base = "https://campusdining.princeton.edu/dining/_Foodpro/"

page = { 'menu':        'menuSamp.asp?',
         'nutr':        'pickMenu.asp?',
         'fact':        'label.asp?' }

dhall = { 'roma':       "locationNum=01",
          'wucox':      "locationNum=02",
          'forbes':     "locationNum=03", 
          'grad':       "locationNum=04", 
          'cjl':        "locationNum=05", 
          'woodywoo':   "locationNum=07", 
          'whitman':    "locationNum=08", 
          'witherspoon':"locationNum=16", 
          'viv':        "locationNum=17", 
          'green':      "locationNum=19", 
          'grill':      "locationNum=21", 
          'chem':       "locationNum=23", 
          'baked':      "locationNum=30", 
          'breakfast':  "locationNum=06", 
          'frist':      "locationNum=15", 
          'salad':      "locationNum=40" }

def scrape(college):
   '''Get menu'''
   f = urllib.urlopen(base + page['menu'] + dhall[college])

   html = f.read()
   soup = BeautifulSoup(html)
   # print soup.prettify()

   meals = {}
   for title in soup.find_all('div', {'id':'menusampmeals'}):
      meal = title.find_parents('table')[1]
      
      foods = []
      for foodname in meal.find_all('div', {'class':'menusamprecipes'}):
         foods.append(getFood(foodname))

      meals[title.text.strip()] = foods

   print meals
   for name in meals:
      getFoodnum(college, name, meals[name])

   return meals


def getFood(foodname):
   ''' Filters '''
   filt = {'Vegan': '#0000FF', 'Vegetarian': '#00FF00', 'Pork': '#8000FF'}

   food = {}
   filts = []

   # check for nuts
   if foodname.text[-1] == 'M':
      filts.append('Nuts')
      food['name'] = foodname.text[:-2]
   else:
      food['name'] = foodname.text
   
   # check for other dietary restrictions
   for key in filt:
      if filt[key] in foodname.span['style']:
         filts.append(key)

   food['filt'] = filts
   return food

def getFoodnum(college, meal, foods):
   
   '''Get menu'''
   f = urllib.urlopen(base + page['nutr'] + dhall[college] + '&meal=' + meal)

   html = f.read()
   soup = BeautifulSoup(html)

   for link in soup.find_all('a', href=True):
      print link['href'].split('&')[-1]

   return 0

def getNutrition(college, )

# https://campusdining.princeton.edu/dining/_Foodpro/label.asp?locationNum=02&RecNumAndPort=020056


def scrapeall():
   foods = {}
   halls = ['roma', 'wucox', 'whitman', 'forbes', 'cjl', 'grad']
   for hall in halls:
      foods[hall] = scrape(hall)
   return foods

if __name__ == "__main__":
  import json
  scrapeall()
  # print json.dumps(scrapeall())







