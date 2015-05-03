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

   meals = {}
   for title in soup.find_all('div', {'id':'menusampmeals'}):
      meal = title.find_parents('table')[1]
      
      foods = []
      for foodname in meal.find_all('div', {'class':'menusamprecipes'}):
         foods.append(getFood(foodname))

      meals[title.text.strip()] = foods

   for name in meals:
      getFoodnum(college, name, meals[name])
      for food in meals[name]:
         if 'num' in food:
            getNutrition(college, food)
         else:
            food['facts'] = {}

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
   f = urllib.urlopen(base + page['nutr'] + dhall[college] + '&mealName=' + meal)

   html = f.read()
   soup = BeautifulSoup(html)

   for link in soup.find_all('a', href=True):
      url = link['href'].split('&')[-1]
      if url.find('RecNum') != -1:
         food = [item for item in foods if item["name"] == link.text]
         if len(food) > 0:
            food[0]['num'] = url


def getNutrition(college, food):
   facts = {}

   '''Get menu'''
   f = urllib.urlopen(base + page['fact'] + dhall[college] + '&' + food['num'])

   html = f.read()
   soup = BeautifulSoup(html)

   print food

   for fact in soup.find_all('div', {'id': 'facts4'}):
      print fact.text



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







