# Book Management System
Developed a online Book Management Systems, where you can add a new book having its name, author, publication & price. This has all http verbs means you can save, delete, update & fetch book.
As of now I developed book management systems in backend & also tested with **Thunder Client** as well as in **Postman** everything is working fine. Later I will integrate with React.
Although I tried to display some content using html.

# Libraries :
- **Frame work** : Gorilla mux framework
- **ORM** : gorm
- **Database** : mysql
- **JSON marshall, unmarshal**

# Component
- Module
- Controller
- Config
- Model
- utils
- Routes

# Package Structor
- cmd
    - main
         - main.go
     
- config
    - app.go
- controllers
    - book-controller.go
- models
    - book.go
- routes
    - bookstore-routes.go
- utils
    - utils.go

# Controllers Functions/Endpoints
# Function                # EndPoints                    #Http verbs
- GetBooks                    /book                          POST

