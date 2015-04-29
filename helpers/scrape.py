def scrape(college,):
   import urllib
   import string
   import re
   from bs4 import BeautifulSoup

   '''URL parameters'''
   base = "https://campusdining.princeton.edu/dining/_Foodpro/menuSamp.asp?%s"

   dhall = { 'roma': "locationNum=01&locationName=Rockefeller+%26+Mathey+Colleges&naFlag=1",
             'wucox': "locationNum=02&locationName=Butler+%26+Wilson+Colleges&naFlag=1",
             'forbes': "locationNum=03&locationName=Forbes+College&naFlag=1", 
             'grad': "locationNum=04&locationName=Graduate+College+&naFlag=1", 
             'cjl': "locationNum=05&locationName=Center+for+Jewish+Life&naFlag=1", 
             'woodywoo': "locationNum=07&locationName=Woodrow+Wilson+Cafe&naFlag=1", 
             'whitman': "locationNum=08&locationName=Whitman+College&naFlag=1", 
             'witherspoon': "locationNum=16&locationName=Witherspoon%27s&naFlag=1", 
             'viv': "locationNum=17&locationName=Cafe+Vivian&naFlag=1", 
             'green': "locationNum=19&locationName=Chancellor+Green+Cafe&naFlag=1", 
             'grill': "locationNum=21&locationName=Every+Day+Grill&naFlag=1", 
             'chem': "locationNum=23&locationName=Chemistry+CaFe&naFlag=1", 
             'baked': "locationNum=30&locationName=Baked+Goods+%26+Frozen+Treats&naFlag=1", 
             'breakfast': "locationNum=06&locationName=Breakfast+%26+Brunch&naFlag=1", 
             'frist': "locationNum=15&locationName=Frist+Gallery&naFlag=1", 
             'salad': "locationNum=40&locationName=Salad+Selections&naFlag=1"}

   ''' Filters '''
   filt = {'Vegan': '#0000FF', 'Vegetarian': '#00FF00', 'Pork': '#8000FF'}

   '''Get menu'''
   f = urllib.urlopen(base % dhall[college])

   html = f.read()
   soup = BeautifulSoup(html)
   '''print soup.prettify()'''

   meals = {}
   for title in soup.find_all('div', {'id':'menusampmeals'}):
      meal = title.find_parents('table')[1]
      
      foods = []
      for foodname in meal.find_all('div', {'class':'menusamprecipes'}):
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

         # set things up
         food['filt'] = filts
         foods.append(food)


         ####### other filters

         # if foodname.find('span', {'class': 'menusampspecialchars'}) is not None:
         #  # nuts

         # for tag in foodname.parent.next_siblings:
         #  print tag # image for special filters


      meals[title.text.strip()] = foods


   return meals

def scrapeall():
   foods = {}
   halls = ['roma', 'wucox', 'whitman', 'forbes', 'grad', 'cjl']
   for hall in halls:
      foods[hall] = scrape(hall)
   return foods

if __name__ == "__main__":
  import json
  print json.dumps(scrapeall())







