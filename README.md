# oapigen
Code to open API specification without describe

Currently i got this TBC
```
❯ go run . -debug -p examples/fiber/
2025-11-15T18:31:24+07:00 DBG Found func=Success
2025-11-15T18:31:24+07:00 DBG Found func=NewBookHandler
2025-11-15T18:31:24+07:00 DBG Found func=NewBookDB
2025-11-15T18:31:24+07:00 DBG Found func=*bookUsecase.Delete
2025-11-15T18:31:24+07:00 DBG Found func=*bookHandler.CreateBook
2025-11-15T18:31:24+07:00 DBG Found func=*bookHandler.GetBook
2025-11-15T18:31:24+07:00 DBG Found func=*bookHandler.DeleteBook
2025-11-15T18:31:24+07:00 DBG Found func=*bookRepository.Update
2025-11-15T18:31:24+07:00 DBG Found func=NewBookUsecase
2025-11-15T18:31:24+07:00 DBG Found func=main
2025-11-15T18:31:24+07:00 DBG Found func=SuccessPagination
2025-11-15T18:31:24+07:00 DBG Found func=*bookHandler.UpdateBook
2025-11-15T18:31:24+07:00 DBG Found func=*bookRepository.Delete
2025-11-15T18:31:24+07:00 DBG Found func=*bookUsecase.Create
2025-11-15T18:31:24+07:00 DBG Found func=*bookUsecase.Update
2025-11-15T18:31:24+07:00 DBG Found func=NewBookRepository
2025-11-15T18:31:24+07:00 DBG Found func=*bookRepository.Create
2025-11-15T18:31:24+07:00 DBG Found func=*bookRepository.Read
2025-11-15T18:31:24+07:00 DBG Found func=*bookUsecase.Read
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/book/dto.go
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/book/handler.go
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/book/model.go
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/book/repository.go
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/book/usecase.go
2025-11-15T18:31:24+07:00 DBG Processing filename=examples/fiber/main.go
2025-11-15T18:31:24+07:00 DBG Processing func=CreateBook
2025-11-15T18:31:24+07:00 DBG Processing func=GetBook
2025-11-15T18:31:24+07:00 DBG Processing func=UpdateBook
2025-11-15T18:31:24+07:00 DBG Processing func=DeleteBook
2025-11-15T18:31:24+07:00 DBG Found route body=CreateBookRequest handler=bookHandler.CreateBook method=post path=/books response=BookResponse
2025-11-15T18:31:24+07:00 DBG Found route body= handler=bookHandler.GetBook method=get path=/books/:id response=BookResponse
2025-11-15T18:31:24+07:00 DBG Found route body=UpdateBookRequest handler=bookHandler.UpdateBook method=put path=/books/:id response=BookResponse
2025-11-15T18:31:24+07:00 DBG Found route body= handler=bookHandler.DeleteBook method=delete path=/books/:id response=
```
