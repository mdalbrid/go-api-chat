# go-api-chat

----
###### REST API
----

### Для запуска приложения:

```
make docker_all && make migrate_up
```


### Запросы

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
