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

# Controllers Functions & Endpoints

|      Functions          |      Endpoints          |    Http verbs
|-------------------------|-------------------------|-------------------------
|  CreateBook             |   /book                 |    POST
|  GetBooks               |   /book                 |    GET
|  GetBookById            |   /book/{bookId}        |    GET
|  UpdateBook             |   /book/{bookId}        |    PUT
|  DeleteBook             |   /book/{bookId}        |    DELETE

# Fork 
- Don't forget to clone this repositroy
- https://github.com/Sufiyan33/goLang.git

- Follow me for more such types projects.
