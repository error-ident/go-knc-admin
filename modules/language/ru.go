// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import "strings"

var ru = LangSet{
	"managers":         "Менеджеры",
	"name":             "Имя",
	"nickname":         "Никнейм",
	"role":             "Роль",
	"createdat":        "Дата создания",
	"updatedat":        "Дата обновления",
	"path":             "Путь",
	"new":              "Новый",
	"filter":           "Фильтр",
	"action":           "Действие",
	"toggle dropdown":  "Раскрыть меню",
	"delete":           "Удалить",
	"refresh":          "Обновить",
	"expand":           "Развернуть",
	"collapse":         "Свернуть",
	"back":             "Назад",
	"reset":            "Сбросить",
	"save":             "Сохранить",
	"edit":             "Редактировать",
	"operation":        "Операция",
	"method":           "Метод",
	"input":            "Входные данные",
	"online":           "Онлайн",
	"setting":          "Настройки",
	"sign out":         "Выйти",
	"all":              "Все",
	"confirm password": "Подтвердите пароль",
	"search":           "Поиск",
	"remove":           "Удалить",

	"goadmin is now running. \nrunning in \"debug\" mode. switch to \"release\" mode in production.\n\n": "GoAdmin запущен. \nРаботает в режиме \"отладки\". В продакшене переключитесь на режим \"релиза\".\n\n",

	"wrong goadmin version, theme %s required goadmin version are %s":    "Неверная версия GoAdmin, для темы %s требуется версия GoAdmin %s",
	"wrong theme version, goadmin %s required version of theme %s is %s": "Неверная версия темы, для GoAdmin %s требуется версия темы %s %s",

	"adapter is nil, import the default adapter or use addadapter method add the adapter": "адаптер не установлен, импортируйте стандартный адаптер или используйте метод AddAdapter для добавления адаптера",

	"are you sure to delete":               "Вы уверены, что хотите удалить",
	"yes":                                  "да",
	"got it":                               "понятно",
	"cancel":                               "отмена",
	"refresh succeeded":                    "Обновление успешно",
	"reload succeeded":                     "Перезагрузка успешна",
	"all method if empty":                  "Все методы, если пусто",
	"password does not match":              "Пароль не совпадает",
	"should be unique":                     "Должно быть уникальным",
	"slug exists":                          "Slug уже существует",
	"no corresponding options?":            "Нет соответствующих опций?",
	"create here.":                         "Создать здесь.",
	"use for login":                        "Использовать для входа",
	"use to display":                       "Использовать для отображения",
	"a path a line, without global prefix": "Путь в одну строку, без глобального префикса",
	"slug or http_path or name should not be empty": "Slug, http_path или имя не должны быть пустыми",
	"no roles":          "нет ролей",
	"fixed the sidebar": "Фиксированный боковой столбец",
	"enter fullscreen":  "Перейти в полноэкранный режим",
	"exit fullscreen":   "Выйти из полноэкранного режима",

	"permission manage": "Управление правами",
	"menus manage":      "Управление меню",
	"roles manage":      "Управление ролями",
	"operation log":     "Журнал операций",

	"continue editing":  "Продолжить редактирование",
	"continue creating": "Продолжить создание",

	"browse":     "Обзор",
	"avatar":     "Аватар",
	"password":   "Пароль",
	"username":   "Имя пользователя",
	"slug":       "Slug",
	"permission": "Права доступа",
	"userid":     "Идентификатор пользователя",
	"content":    "Содержание",
	"parent":     "Родительский элемент",
	"icon":       "Иконка",
	"uri":        "Uri",

	"detail": "Подробности",

	"admin":     "Администратор",
	"users":     "Пользователи",
	"roles":     "Роли",
	"menu":      "Меню",
	"dashboard": "Панель управления",
	"home":      "Домой",

	"initialize configuration":        "Инициализация конфигурации",
	"initialize navigation buttons":   "Инициализация кнопок навигации",
	"initialize plugins":              "Инициализация плагинов",
	"initialize database connections": "Инициализация подключений к базе данных",
	"initialize success":              "Инициализация успешно🍺🍺",

	"not found":      "Не найдено",
	"internal error": "Внутренняя ошибка",
	"unauthorized":   "Неавторизованный",
	"plugin setting": "Настройки плагина",

	"plugins":          "Плагины",
	"plugin store":     "Магазин плагинов",
	"get more plugins": "Получить больше плагинов",
	"uninstalled":      "Удален",

	"second":  "секунда",
	"seconds": "секунды",
	"minute":  "минута",
	"minutes": "минуты",
	"hour":    "час",
	"hours":   "часы",
	"day":     "день",
	"days":    "дни",
	"week":    "неделя",
	"weeks":   "недели",
	"month":   "месяц",
	"months":  "месяцы",
	"year":    "год",
	"years":   "годы",

	"config.domain":          "Домен сайта",
	"config.language":        "Язык сайта",
	"config.url prefix":      "URL-префикс",
	"config.theme":           "Тема",
	"config.title":           "Заголовок",
	"config.index url":       "URL главной страницы",
	"config.login url":       "URL страницы входа",
	"config.env":             "Среда",
	"config.color scheme":    "Цветовая схема",
	"config.cdn url":         "URL к CDN-хранилищу ресурсов",
	"config.login title":     "Заголовок страницы входа",
	"config.auth user table": "Таблица пользователей",
	"config.extra":           "Дополнительная конфигурация",
	"config.store":           "Настройки файлового хранилища",
	"config.databases":       "Настройки базы данных",

	"config.general":      "Общие настройки",
	"config.log":          "Лог",
	"config.site setting": "Настройки сайта",
	"config.custom":       "Настраиваемые настройки",
	"config.debug":        "Режим отладки",
	"config.site off":     "Сайт недоступен",
	"config.true":         "Включено",
	"config.false":        "Выключено",

	"config.logo":                        "Логотип",
	"config.mini logo":                   "Мини-логотип",
	"config.bootstrap file path":         "Путь к файлу Bootstrap",
	"config.go mod file path":            "Путь к файлу go.mod",
	"config.session life time":           "Время жизни сессии",
	"config.custom head html":            "HTML заголовка",
	"config.custom foot html":            "HTML подвала",
	"config.custom 404 html":             "404 страница",
	"config.custom 403 html":             "403 страница",
	"config.custom 500 html":             "500 страница",
	"config.hide config center entrance": "Скрыть кнопку конфигурации",
	"config.hide app info entrance":      "Скрыть кнопку информации о приложении",
	"config.hide tool entrance":          "Скрыть кнопку инструментов",
	"config.footer info":                 "Информация в подвале",
	"config.login logo":                  "Логотип при входе в систему",
	"config.no limit login ip":           "Неограниченный вход с нескольких IP-адресов",
	"config.operation log off":           "Отключить журнал операций",
	"config.allow delete operation log":  "Разрешить удаление журнала операций",
	"config.animation type":              "Тип анимации",
	"config.animation duration":          "Длительность анимации (сек)",
	"config.animation delay":             "Задержка анимации (сек)",
	"config.file upload engine":          "Движок загрузки файлов",

	"config.logger rotate":             "Настройки поворота логов",
	"config.logger rotate max size":    "Макс. размер (мб)",
	"config.logger rotate max backups": "Макс. количество резервных копий",
	"config.logger rotate max age":     "Макс. возраст (дней)",
	"config.logger rotate compress":    "Сжатие",

	"config.info log path":         "Путь к файлу журнала информации",
	"config.error log path":        "Путь к файлу журнала ошибок",
	"config.access log path":       "Путь к файлу журнала доступа",
	"config.info log off":          "Выключить журнал информации",
	"config.error log off":         "Выключить журнал ошибок",
	"config.access log off":        "Выключить журнал доступа",
	"config.access assets log off": "Выключить журнал доступа к ресурсам",
	"config.sql log on":            "Включить журнал SQL",
	"config.log level":             "Уровень журнала",

	"config.logger rotate encoder":                "Настройки кодирования журнала",
	"config.logger rotate encoder time key":       "Ключ времени",
	"config.logger rotate encoder level key":      "Ключ уровня",
	"config.logger rotate encoder name key":       "Ключ имени",
	"config.logger rotate encoder caller key":     "Ключ вызывающей функции",
	"config.logger rotate encoder message key":    "Ключ сообщения",
	"config.logger rotate encoder stacktrace key": "Ключ стека вызовов",
	"config.logger rotate encoder level":          "Кодировщик уровня",
	"config.logger rotate encoder time":           "Кодировщик времени",
	"config.logger rotate encoder duration":       "Кодировщик продолжительности",
	"config.logger rotate encoder caller":         "Кодировщик вызывающей функции",
	"config.logger rotate encoder encoding":       "Формат вывода",

	"config.capital":        "Заглавные буквы",
	"config.capitalcolor":   "Заглавные буквы с цветом",
	"config.lowercase":      "Строчные буквы",
	"config.lowercasecolor": "Строчные буквы с цветом",

	"config.seconds":     "Секунды",
	"config.nanosecond":  "Наносекунды",
	"config.microsecond": "Микросекунды",
	"config.millisecond": "Миллисекунды",

	"config.full path":  "Полный путь",
	"config.short path": "Короткий путь",
	//

	"config.do not modify when you have not set up all assets": "Не изменяйте, когда не настроены все ресурсы",
	"config.it will work when theme is adminlte":               "Работает только с темой AdminLTE",

	"config.language." + CN:                  "Китайский",
	"config.language." + EN:                  "Английский",
	"config.language." + JP:                  "Японский",
	"config.language." + strings.ToLower(TC): "Традиционный китайский",
	"config.language." + PTBR:                "Бразильский португальский",
	"config.language." + RU:                  "Русский",

	"config.modify site config":         "Изменение конфигурации сайта",
	"config.modify site config success": "Успешно изменено",
	"config.modify site config fail":    "Не удалось изменить",

	"system.system info":     "Информация о системе и приложении",
	"system.application":     "Информация о приложении",
	"system.application run": "Информация о запущенных приложениях",
	"system.system":          "Информация о системе",

	"system.process_id":                           "ID процесса",
	"system.golang_version":                       "Версия Golang",
	"system.server_uptime":                        "Время работы сервера",
	"system.current_goroutine":                    "Текущие горутины",
	"system.current_memory_usage":                 "Использование памяти",
	"system.total_memory_allocated":               "Общее количество выделенной памяти",
	"system.memory_obtained":                      "Полученная память",
	"system.pointer_lookup_times":                 "Время поиска указателя",
	"system.memory_allocate_times":                "Время выделения памяти",
	"system.memory_free_times":                    "Время освобождения памяти",
	"system.current_heap_usage":                   "Использование кучи",
	"system.heap_memory_obtained":                 "Полученная куча памяти",
	"system.heap_memory_idle":                     "Неиспользуемая куча памяти",
	"system.heap_memory_in_use":                   "Используемая куча памяти",
	"system.heap_memory_released":                 "Освобожденная куча памяти",
	"system.heap_objects":                         "Объекты кучи",
	"system.bootstrap_stack_usage":                "Использование стека Bootstrap",
	"system.stack_memory_obtained":                "Полученная память стека",
	"system.mspan_structures_usage":               "Использование структур MSpan",
	"system.mspan_structures_obtained":            "Полученные структуры MSpan",
	"system.mcache_structures_usage":              "Использование структур MCache",
	"system.mcache_structures_obtained":           "Полученные структуры MCache",
	"system.profiling_bucket_hash_table_obtained": "Полученная таблица хешей ведра профилирования",
	"system.gc_metadata_obtained":                 "Полученная метадата GC",
	"system.other_system_allocation_obtained":     "Другие системные выделения",
	"system.next_gc_recycle":                      "Следующая переработка GC",
	"system.last_gc_time":                         "Время с момента последней переработки GC",
	"system.total_gc_time":                        "Общая приостановка GC",
	"system.total_gc_pause":                       "Общая приостановка GC",
	"system.last_gc_pause":                        "Последняя приостановка GC",
	"system.gc_times":                             "Количество переработок GC",

	"system.cpu_logical_core": "Логические ядра CPU",
	"system.cpu_core":         "Физические ядра CPU",
	"system.os_platform":      "Платформа ОС",
	"system.os_family":        "Семейство ОС",
	"system.os_version":       "Версия ОС",
	"system.load1":            "Загрузка 1",
	"system.load5":            "Загрузка 5",
	"system.load15":           "Загрузка 15",
	"system.mem_total":        "Общий объем памяти",
	"system.mem_available":    "Доступный объем памяти",
	"system.mem_used":         "Использованный объем памяти",

	"system.app_name":         "Название приложения",
	"system.go_admin_version": "Версия приложения",
	"system.theme_name":       "Тема",
	"system.theme_version":    "Версия темы",

	"tool.tool":                 "Инструмент",
	"tool.table":                "Таблица",
	"tool.connection":           "Подключение",
	"tool.package":              "Пакет",
	"tool.output":               "Путь вывода",
	"tool.output path is empty": "Путь вывода пуст",
	"tool.field":                "Поле",
	"tool.title":                "Название",
	"tool.field name":           "Имя поля",
	"tool.db type":              "Тип базы данных",
	"tool.form type":            "Тип формы",
	"tool.generate table model": "Создать модель таблицы",
	"tool.primarykey":           "Первичный ключ",
	"tool.field filterable":     "Фильтруемое",
	"tool.field sortable":       "Сортируемое",
	"tool.yes":                  "Да",
	"tool.no":                   "Нет",
	"tool.hide":                 "Скрыть",
	"tool.show":                 "Показать",
	"tool.generate success":     "Успешно создано",
	"tool.hide filter area":     "Скрыть область фильтров",
	"tool.use absolute path":    "Использовать абсолютный путь",
	"tool.display":              "Отображение",
	"tool.basic info":           "Основная информация",
	"tool.table info":           "Информация о таблице",
	"tool.form info":            "Информация о форме",
	"tool.field editable":       "Редактируемое",
	"tool.info field editable":  "Редактируемое",
	"tool.extra import package": "Дополнительный пакет импорта",
	"tool.field can add":        "Можно добавлять",
	"tool.field default":        "Значение по умолчанию",
	"tool.filter area":          "Область фильтрации",
	"tool.new button":           "Кнопка Новый",
	"tool.export button":        "Кнопка Экспорт",
	"tool.edit button":          "Кнопка Редактировать",
	"tool.delete button":        "Кнопка Удалить",
	"tool.detail button":        "Кнопка Подробно",
	"tool.filter button":        "Кнопка Фильтр",
	"tool.row selector":         "Селектор строки",
	"tool.pagination":           "Пагинация",
	"tool.query info":           "Информация о запросе",
	"tool.filter form layout":   "Макет формы фильтрации",
	//
	"tool.continue edit checkbox": "Флажок продолжить редактирование",
	"tool.continue new checkbox":  "Флажок продолжить создание",
	"tool.reset button":           "Кнопка сброса",
	"tool.back button":            "Кнопка назад",
	"tool.generate":               "Сгенерировать",
	"tool.generated tables":       "Сгенерированные таблицы",
	"tool.description":            "Описание",
	"tool.label":                  "Метка",
	"tool.image":                  "Изображение",
	"tool.bool":                   "Логическое значение",
	"tool.link":                   "Ссылка",
	"tool.fileSize":               "Размер файла",
	"tool.date":                   "Дата",
	"tool.icon":                   "Иконка",
	"tool.dot":                    "Точка",
	"tool.progressBar":            "Индикатор выполнения",
	"tool.loading":                "Загрузка",
	"tool.downLoadable":           "Можно скачать",
	"tool.copyable":               "Можно копировать",
	"tool.carousel":               "Карусель",
	"tool.qrcode":                 "QR-код",
	"tool.field hide":             "Скрыть поле",
	"tool.field display":          "Отобразить поле",
	"tool.table permission":       "Сгенерировать права доступа",
	"tool.extra code":             "Дополнительный код",
	//
	"tool.field display normal":     "Обычный",
	"tool.field diplay hide":        "Скрыт",
	"tool.field diplay edit hide":   "Скрыть при редактировании",
	"tool.field diplay create hide": "Скрыть при создании",

	"tool.generate table model success": "Успешно сгенерирована модель таблицы",
	"tool.generate table model fail":    "Не удалось сгенерировать модель таблицы",

	"tool.detail display":             "Отображение деталей",
	"tool.detail info":                "Информация о деталях",
	"tool.follow list page":           "Следовать за страницей списка",
	"tool.inherit from list page":     "Наследовать от страницы списка",
	"tool.independent from list page": "Независимо от страницы списка",

	"generator.query":                 "Запрос",
	"generator.show edit form page":   "Показать страницу формы редактирования",
	"generator.show create form page": "Показать страницу формы создания",
	"generator.edit":                  "Редактировать",
	"generator.create":                "Создать",
	"generator.delete":                "Удалить",
	"generator.export":                "Экспортировать",

	"plugin.plugin":                         "Плагин",
	"plugin.plugin detail":                  "Детали плагина",
	"plugin.introduction":                   "Введение",
	"plugin.website":                        "Веб-сайт",
	"plugin.version":                        "Версия",
	"plugin.created at":                     "Дата создания",
	"plugin.updated at":                     "Дата обновления",
	"plugin.provided by %s":                 "Предоставлено %s",
	"plugin.upgrade":                        "Обновить",
	"plugin.install":                        "Установить",
	"plugin.download":                       "Скачать",
	"plugin.buy":                            "Купить",
	"plugin.downloading":                    "Загрузка",
	"plugin.info":                           "Информация",
	"plugin.login":                          "Вход",
	"plugin.login to goadmin member system": "Войти в систему участников GoAdmin",
	"plugin.account":                        "Аккаунт",
	"plugin.password":                       "Пароль",
	"plugin.learn more":                     "Узнать больше",

	"plugin.no account? click %s here %s to register.":    "Нет аккаунта? Нажмите %sздесь%s, чтобы зарегистрироваться.",
	"plugin.download fail, wrong name":                    "Сбой загрузки, неправильное имя",
	"plugin.change to debug mode first":                   "Сначала переключитесь в режим отладки",
	"plugin.download fail, plugin not exist":              "Сбой загрузки, плагин не существует",
	"plugin.download fail":                                "Сбой загрузки",
	"plugin.golang develop environment does not exist":    "Среда разработки Golang не существует",
	"plugin.download success, restart to install":         "Успешная загрузка, перезапустите для установки",
	"plugin.restart to install":                           "Перезапустите для установки",
	"plugin.can not connect to the goadmin remote server": "Не удалось подключиться к удаленному серверу GoAdmin, проверьте соединение с сетью.",

	"admin.basic admin": "Базовый администратор",
	"admin.a built-in plugins of goadmin which help you to build a crud manager platform quickly.": "Встроенные плагины GoAdmin, которые помогают вам быстро создавать платформу для управления crud.",
	"admin.official": "Официальный",
}