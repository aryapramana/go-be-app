# go-be-app
**Middleware App - API Station With Go**


Pre-requisite:
1. Go to commonfile\constantconfig.go and replace the api_key with your personal key
2. Go to dbconnect\conndb.go and replace the connection with your personal db
3. Setup your db (screenshot provided)


API station with Go. 
Feature:
1. Post - middle (json post httpbin.org)
2. Post - db (json post direct db)
3. Get - middle (get data from newsapi.org with country selected)
4. Get - db (get data from direct db)
5. Get - db (get data from direct db with query)
6. Put - db (update data from direct db)
7. Delete - db (delete data from direct db)



**DB Structure**
<img width="929" alt="db-structure" src="https://user-images.githubusercontent.com/40905337/156918922-0b587a1e-0c73-4fbf-b591-8a95aa0b85c7.png">

**Screenshot**
delete data
<img width="1047" alt="delete-data-db" src="https://user-images.githubusercontent.com/40905337/156918955-62ae93dd-afab-4d31-b21d-437e5dffd550.png">

get data db
<img width="1057" alt="get-data-db" src="https://user-images.githubusercontent.com/40905337/156918967-c3cdca71-a64c-4f83-9455-51dda495b401.png">
<img width="1062" alt="get-data-db-query" src="https://user-images.githubusercontent.com/40905337/156918975-5876b139-00dc-4fa7-bb17-c994da7bda3a.png">
<img width="1062" alt="news-mid-getall" src="https://user-images.githubusercontent.com/40905337/156918979-14af045a-07c7-439d-aab9-799e4776fb1d.png">

post data
<img width="1063" alt="post-data-db" src="https://user-images.githubusercontent.com/40905337/156918987-492d5d39-6692-474b-a8d5-acdeaec4e230.png">
<img width="1056" alt="post-mid-bin" src="https://user-images.githubusercontent.com/40905337/156918989-103d64bc-701d-47ee-ad22-8ac594334943.png">

update data
<img width="1056" alt="update-data-db" src="https://user-images.githubusercontent.com/40905337/156919001-41166e87-ea17-4154-86a2-028178e7e025.png">
