## gql study

go mod init xxx

go run github.com/99designs/gqlgen init

mdf gqlgen.yml

go run github.com/99designs/gqlgen generate

go run server.go

### ext

#### split schema before and after

```
mutation createItem {
  createItem(input: {name: "item01"})
}
```
 
 ```{
      "data": {
        "createItem": "item-id-01"
      }
    }
```

```
mutation createUser {
  createUser(input: {name: "user01"})
}
```

```
{
  "data": {
    "createUser": "user-id-01"
  }
}
```

```
query item  {
  item(id: "item-id-01") {
    id
    name
  }
}
```

```
{
  "data": {
    "item": {
      "id": "item-id-01",
      "name": "item01"
    }
  }
}
```

```
query user  {
  user(id: "user-id-01") {
    id
    name
  }
}
```

```
{
  "data": {
    "user": {
      "id": "user-id-01",
      "name": "user01"
    }
  }
}
```