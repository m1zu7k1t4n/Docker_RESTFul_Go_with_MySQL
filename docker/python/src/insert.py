import mysql.connector
import random

connector = mysql.connector.connect(
    user='test', password='test', host='mysql_host', database='test')

cursor = connector.cursor()
try:
  for _ in range(10000):
    cursor.execute("INSERT INTO tasks(title, body) VALUES({0},{1})".format(
        random.randint(10000, 10000000), random.randint(10000, 10000000)))
  connector.commit()
except:
  connector.rollback()
  raise
cursor.close()

connector.close()
