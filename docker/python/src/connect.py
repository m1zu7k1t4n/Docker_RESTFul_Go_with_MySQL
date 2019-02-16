import mysql.connector

connector = mysql.connector.connect(user='test', password='test', host='mysql_host', database='test')

cursor = connector.cursor(dictionary=True)
cursor.execute("select * from tasks;")
for row in cursor.fetchall()[:50]:
    print(row)
cursor.close()

connector.close()
