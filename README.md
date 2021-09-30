# dam.example_app

Проект содержит примеры и материалы для обучения работе с DAM.

## По проекту

В корне проекта есть каталоги `mongodb`, `my-app`, `my-resources`.
Они являются примерами приложений DAM.

При работе с `my-app` необходимо предварительно создать бинарный `my-app` файл.
Для этого используется Go версии 1.16 и выше.

В проекте есть текстовые файлы `ЧастьX.md` с описанием, как работать с примерами и DAM.
Части 1-6 были подготовлены для видео на [youtube](https://www.youtube.com/playlist?list=PLCMW4HaNNZOUK8pOFVydvWlxIRi0bTISt).
После публикации данному проекту назначен тэг [Youtube-Video-Release](https://github.com/dam-utils/dam.example_app/tree/Youtube-Video-Release)

Примеры для частей 1-6 запускаются в ветке `main` на Ubuntu 20.04.

В ветке `windows` примеры дорабатываются под Windows 10 x64.

## Для чего DAM ?

DAM в GitLab [https://github.com/dam-utils/dam](https://github.com/dam-utils/dam)

* облегчает разработку и тестирование микро сервисных приложений,
  для локального развертывания без привязки к среде разработки,
  а также архитектуре операционной системы

* для развертывания и поддержке продуктов на машинах с ограниченным доступом к интернет

* администраторам и сопровождению, чтобы облегчить дальнейшую настройку и контролировать обновление продукта

## Статья, описание к курсу видео по созданию приложений

- [Часть 1. Что такое DAM](Часть1.md)
- [Часть 2. Создание приложений](Часть2.md)
- [Часть 3. Работа с продуктами](Часть3.md)
- [Часть 4. Поиск. Работа с репозиториями](Часть4.md)
- [Часть 5. Метки. Чистка системы](Часть5.md)
- [Часть 6. Сборка. Конфиг. Разработчикам](Часть6.md)


## dam help
```
Docker Application Manager

Version:
1.2.1

Usage:
dam [command]

Available Commands:
addrepo     Add an app registry to the system.
completion  generate the autocompletion script for the specified shell
create      Create docker application.
export      Export apps to file.
help        Help about any command
import      Import apps from file.
info        Information for your application.
install     Install docker application from a docker registry or a specific file archive.
list        List all installed your applications.
listrepos   List all defined repositories.
modifyrepo  Modify properties of repositories specified.
purge       Remove docker images not attached to apps.
remove      Remove docker application.
removerepo  Remove registry specified by name or number.
save        Save app to an archive.
search      Search app in remote registry.

Flags:
-x, --debug   Enable debug mode
-h, --help    help for dam

Use "dam [command] --help" for more information about a command.
```
