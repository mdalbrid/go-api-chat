# go-api-chat

----
###### REST API
----

### Для запуска приложения:

```
make docker_all && make migrate_up
```


### Запросы

[chat.postman_collection.json](https://github.com/mdalbrid/go-api-chat/blob/main/chat.postman_collection.json "export from Postman") - Postman Export Collection

`/api/chat/`⠀- ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀&nbsp;⠀⠀&nbsp;⠀⠀⠀⠀POST<br>
`/api/chat/{chat_id}/msg/` ⠀-⠀⠀⠀⠀⠀⠀ POST<br>
`/api/chat/{chat_id}/msg/?limit=№` - ⠀GET<br>
`/api/chat/{chat_id}/msg/{msg_id}` - ⠀GET<br>



<br>

----

### TODO:

- сделать авторизацию JWT
- дописать документацию
- сделать инструкцию по настройке
- ¯\ _ (ツ) _ /¯
- 
