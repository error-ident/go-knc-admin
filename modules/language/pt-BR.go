// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import "strings"

var ptbr = LangSet{
	"managers":         "Gerentes",
	"name":             "Nome",
	"nickname":         "Apelido",
	"role":             "Role",
	"createdat":        "criadoEm",
	"updatedat":        "atualizadoEm",
	"path":             "caminho",
	"new":              "Novo",
	"filter":           "Filtro",
	"action":           "Ação",
	"toggle dropdown":  "Toggle Dropdown",
	"delete":           "Excluir",
	"refresh":          "Atualizar",
	"expand":           "Expandir",
	"collapse":         "Recolher",
	"back":             "Voltar",
	"reset":            "Redefinir",
	"save":             "Salvar",
	"edit":             "Editar",
	"operation":        "Operação",
	"method":           "Método",
	"input":            "entrada",
	"online":           "Online",
	"setting":          "Configuração",
	"sign out":         "Sair",
	"all":              "Tudo",
	"confirm password": "Confirmar a Senha",
	"search":           "Procurar",
	"remove":           "Remover",

	"goadmin is now running. \nrunning in \"debug\" mode. switch to \"release\" mode in production.\n\n": "GoAdmin agora está em execução. \nExecutando no modo \"debug \". Mude para o modo de \"release\" em produção. \n\n",

	"wrong goadmin version, theme %s required goadmin version are %s":    "wrong GoAdmin version, theme %s required GoAdmin version are %s",
	"wrong theme version, goadmin %s required version of theme %s is %s": "wrong Theme version, GoAdmin %s required version of theme %s is %s",

	"adapter is nil, import the default adapter or use addadapter method add the adapter": "adapter is nil, import the default adapter or use AddAdapter method add the adapter",

	"are you sure to delete":               "Você tem certeza que quer deletar",
	"yes":                                  "sim",
	"got it":                               "entendi",
	"cancel":                               "cancelar",
	"refresh succeeded":                    "Atualizado com sucesso",
	"reload succeeded":                     "Recarregado com sucesso",
	"all method if empty":                  "Todo método é vazio",
	"password does not match":              "Senha não confere",
	"should be unique":                     "Deve ser único",
	"slug exists":                          "Slug existe",
	"no corresponding options?":            "Sem opções correspondentes ?",
	"create here.":                         "Crie aqui.",
	"use for login":                        "Use para login",
	"use to display":                       "Use para exibir",
	"a path a line, without global prefix": "Um caminho uma linha",
	"slug or http_path or name should not be empty": "slug ou http_path ou nome não deve estar vazio",
	"no roles":          "sem roles",
	"fixed the sidebar": "Barra lateral fixada",
	"enter fullscreen":  "Entrar em tela cheia",
	"exit fullscreen":   "Sair da tela inteira",

	"permission manage": "Gerenciamento de permissão",
	"menus manage":      "Gerenciar menu",
	"roles manage":      "Gerenciar roles",
	"operation log":     "Log de operação",

	"continue editing":  "Continue editando",
	"continue creating": "Continue criando",

	"browse":     "Navegar",
	"avatar":     "Avatar",
	"password":   "Senha",
	"username":   "Nome do usuário",
	"slug":       "Slug",
	"permission": "Permissão",
	"userid":     "UserID",
	"content":    "Conteúdo",
	"parent":     "Pai",
	"icon":       "Ícone",
	"uri":        "Uri",

	"detail": "Detalhe",

	"admin":     "Admin",
	"users":     "Users",
	"roles":     "Roles",
	"menu":      "Menu",
	"dashboard": "Dashboard",
	"home":      "Home",

	"initialize configuration":        "Inicializando configuração",
	"initialize navigation buttons":   "Inicializando botões de navegação",
	"initialize plugins":              "Inicializando plugins",
	"initialize database connections": "Inicializando conexões de banco de dados",
	"initialize success":              "Inicialização feita com sucesso🍺🍺",

	"not found":      "Não encontrado",
	"internal error": "Erro interno",
	"unauthorized":   "Não autorizado",
	"plugin setting": "Configuração de Plugin",

	"plugins":          "Plugins",
	"plugin store":     "Loja de Plugins",
	"get more plugins": "Mais plugins",
	"uninstalled":      "Desinstalado",

	"second":  "segundo",
	"seconds": "segundos",
	"minute":  "minuto",
	"minutes": "minutos",
	"hour":    "hora",
	"hours":   "horass",
	"day":     "dia",
	"days":    "dias",
	"week":    "semana",
	"weeks":   "semanas",
	"month":   "mês",
	"months":  "meses",
	"year":    "ano",
	"years":   "anos",

	"config.domain":          "Domínio do site",
	"config.language":        "Idioma do site",
	"config.url prefix":      "URL Prefix",
	"config.theme":           "Tema",
	"config.title":           "Título",
	"config.index url":       "Home URL",
	"config.login url":       "Login URL",
	"config.env":             "Env",
	"config.color scheme":    "Color Scheme",
	"config.cdn url":         "CDN Asset URL",
	"config.login title":     "Título de Login",
	"config.auth user table": "Auth User Table",
	"config.extra":           "Extra Configuration",
	"config.store":           "File Store Setting",
	"config.databases":       "Database Setting",

	"config.general":      "General",
	"config.log":          "Log",
	"config.site setting": "Site Settings",
	"config.custom":       "Customize",
	"config.debug":        "Debug Mode",
	"config.site off":     "Site Offline",
	"config.true":         "On",
	"config.false":        "Off",

	"config.logo":                        "Logo",
	"config.mini logo":                   "Mini Logo",
	"config.bootstrap file path":         "Bootstrap File Path",
	"config.go mod file path":            "go.mod File Path",
	"config.session life time":           "Session Life Time",
	"config.custom head html":            "Head HTML",
	"config.custom foot html":            "Foot HTML",
	"config.custom 404 html":             "404 Page",
	"config.custom 403 html":             "403 Page",
	"config.custom 500 html":             "500 Page",
	"config.hide config center entrance": "Hide Config Button",
	"config.hide app info entrance":      "Hide App Info Button",
	"config.hide tool entrance":          "Hide Tool Button",
	"config.footer info":                 "Footer Info",
	"config.login logo":                  "Login Logo",
	"config.no limit login ip":           "No Limit Login Multi IPs",
	"config.operation log off":           "Operation Log Off",
	"config.allow delete operation log":  "Allow Delete Operation Log",
	"config.animation type":              "Animation Type",
	"config.animation duration":          "Animation Duration(s)",
	"config.animation delay":             "Animation Delay(s)",
	"config.file upload engine":          "File Upload Engine",

	"config.logger rotate":             "Log Rotate Settings",
	"config.logger rotate max size":    "Max Size（m）",
	"config.logger rotate max backups": "Max Buckups",
	"config.logger rotate max age":     "Max Age（day）",
	"config.logger rotate compress":    "Compress",

	"config.info log path":         "Info Log File Path",
	"config.error log path":        "Error Log File Path",
	"config.access log path":       "Access Log File Path",
	"config.info log off":          "Info Log Off",
	"config.error log off":         "Error Log Off",
	"config.access log off":        "Access Log Off",
	"config.access assets log off": "Access Assets Log Off",
	"config.sql log on":            "Abrir log SQL",
	"config.log level":             "Nível",

	"config.logger rotate encoder":                "Configurações do codificador de registro",
	"config.logger rotate encoder time key":       "Chave de tempo",
	"config.logger rotate encoder level key":      "Chave de Nível",
	"config.logger rotate encoder name key":       "Chave do Nome",
	"config.logger rotate encoder caller key":     "Tecla do chamador",
	"config.logger rotate encoder message key":    "Chave de mensagem",
	"config.logger rotate encoder stacktrace key": "Chave Stacktrace",
	"config.logger rotate encoder level":          "Codificador de Nível",
	"config.logger rotate encoder time":           "Codificador de tempo",
	"config.logger rotate encoder duration":       "Codificador de duração",
	"config.logger rotate encoder caller":         "Codificador de Chamadas",
	"config.logger rotate encoder encoding":       "Formato de saída",

	"config.capital":        "Maiúsculo",
	"config.capitalcolor":   "Maiúsculo com cor",
	"config.lowercase":      "Minúsculas",
	"config.lowercasecolor": "Minúsculas com cor",

	"config.seconds":     "Segundos",
	"config.nanosecond":  "Nanossegundo",
	"config.microsecond": "Microssegundo",
	"config.millisecond": "Milissegundo",

	"config.full path":  "Caminho completo",
	"config.short path": "Caminho curto",

	"config.do not modify when you have not set up all assets": "Não modifique quando você não configurou todos os assets",
	"config.it will work when theme is adminlte":               "Funcionará quando o tema for adminlte",

	"config.language." + CN:                  "Chinese",
	"config.language." + EN:                  "English",
	"config.language." + JP:                  "Japanese",
	"config.language." + strings.ToLower(TC): "Traditional Chinese",
	"config.language." + PTBR:                "Brazilian Portuguese",
	"config.language." + RU:                  "Русский",

	"config.modify site config":         "Modificação da configuração do site",
	"config.modify site config success": "modificado com successo",
	"config.modify site config fail":    "modificação falhou",

	"system.system info":     "Informações do sistema e do aplicativo",
	"system.application":     "Informações do aplicativo",
	"system.application run": "Informações de execução de aplicativos",
	"system.system":          "Informação do sistema",

	"system.process_id":                           "Process ID",
	"system.golang_version":                       "Golang Version",
	"system.server_uptime":                        "Server Uptime",
	"system.current_goroutine":                    "Current Goroutines",
	"system.current_memory_usage":                 "Current Memory Usage",
	"system.total_memory_allocated":               "Total Memory Allocated",
	"system.memory_obtained":                      "Memory Obtained",
	"system.pointer_lookup_times":                 "Pointer Lookup Times",
	"system.memory_allocate_times":                "Memory Allocate Times",
	"system.memory_free_times":                    "Memory Free Times",
	"system.current_heap_usage":                   "Current Heap Usage",
	"system.heap_memory_obtained":                 "Heap Memory Obtained",
	"system.heap_memory_idle":                     "Heap Memory Idle",
	"system.heap_memory_in_use":                   "Heap Memory In Use",
	"system.heap_memory_released":                 "Heap Memory Released",
	"system.heap_objects":                         "Heap Objects",
	"system.bootstrap_stack_usage":                "Bootstrap Stack Usage",
	"system.stack_memory_obtained":                "Stack Memory Obtained",
	"system.mspan_structures_usage":               "MSpan Structures Usage",
	"system.mspan_structures_obtained":            "MSpan Structures Obtained",
	"system.mcache_structures_usage":              "MCache Structures Usage",
	"system.mcache_structures_obtained":           "MCache Structures Obtained",
	"system.profiling_bucket_hash_table_obtained": "Profiling Bucket Hash Table Obtained",
	"system.gc_metadata_obtained":                 "GC Metadata Obtained",
	"system.other_system_allocation_obtained":     "Other System Allocation Obtained",
	"system.next_gc_recycle":                      "Next GC Recycle",
	"system.last_gc_time":                         "Since Last GC Time",
	"system.total_gc_time":                        "Total GC Pause",
	"system.total_gc_pause":                       "Total GC Pause",
	"system.last_gc_pause":                        "Last GC Pause",
	"system.gc_times":                             "GC Times",

	"system.cpu_logical_core": "CPU Core Lógico",
	"system.cpu_core":         "CPU Core Físico",
	"system.os_platform":      "OS Plataforma",
	"system.os_family":        "OS Familia",
	"system.os_version":       "OS Versão",
	"system.load1":            "Load1",
	"system.load5":            "Load5",
	"system.load15":           "Load15",
	"system.mem_total":        "Memória Total",
	"system.mem_available":    "Memória disponível",
	"system.mem_used":         "Memória usada",

	"system.app_name":         "Nome da aplicação",
	"system.go_admin_version": "Versão da aplicação",
	"system.theme_name":       "Tema",
	"system.theme_version":    "Versão do Tema",

	"tool.tool":                   "Ferramenta",
	"tool.table":                  "Tabela",
	"tool.connection":             "Conexão",
	"tool.package":                "Pacote",
	"tool.output":                 "Caminho de saída",
	"tool.output path is empty":   "O caminho de saída está vazio",
	"tool.field":                  "Campo",
	"tool.title":                  "Título",
	"tool.field name":             "Nome",
	"tool.db type":                "Tipo de banco de dados",
	"tool.form type":              "Tipo de Formulário",
	"tool.generate table model":   "Modelo de tabela gerado",
	"tool.primarykey":             "Chave primária",
	"tool.field filterable":       "Filtrável",
	"tool.field sortable":         "Classificável",
	"tool.yes":                    "Sim",
	"tool.no":                     "Não",
	"tool.hide":                   "Esconder",
	"tool.show":                   "Mostrar",
	"tool.generate success":       "Gerado com sucesso",
	"tool.hide filter area":       "Esconder a Área do Filtro",
	"tool.use absolute path":      "Use o caminho absoluto",
	"tool.display":                "Exibição",
	"tool.basic info":             "Básico",
	"tool.table info":             "Tabela",
	"tool.form info":              "Formulário",
	"tool.field editable":         "Editável",
	"tool.info field editable":    "Editável",
	"tool.extra import package":   "Pacote de Importação",
	"tool.field can add":          "Pode adicionar",
	"tool.field default":          "Default",
	"tool.filter area":            "Area de filtro",
	"tool.new button":             "Botão Novo",
	"tool.export button":          "Botão exportar",
	"tool.edit button":            "Botão editar",
	"tool.delete button":          "Botão excluir",
	"tool.detail button":          "Botão Detalhe",
	"tool.filter button":          "Botão Filtro",
	"tool.row selector":           "Seletor de linha",
	"tool.pagination":             "Paginação",
	"tool.query info":             "Informações de consulta",
	"tool.filter form layout":     "Layout do formulário de filtro",
	"tool.continue edit checkbox": "Continue Edição Checkbox",
	"tool.continue new checkbox":  "Continuar Novo Checkbox",
	"tool.reset button":           "Botão Redefinir",
	"tool.back button":            "Botão Voltar",
	"tool.generate":               "Gerado",
	"tool.generated tables":       "Tabelas geradas",
	"tool.description":            "Descrição",
	"tool.label":                  "Etiqueta",
	"tool.image":                  "Imagem",
	"tool.bool":                   "Bool",
	"tool.link":                   "Link",
	"tool.fileSize":               "Tamanho do Arquivo",
	"tool.date":                   "Data",
	"tool.icon":                   "Ícone",
	"tool.dot":                    "Ponto",
	"tool.progressBar":            "Barra de progresso",
	"tool.loading":                "Carregando",
	"tool.downLoadable":           "DownLoadable",
	"tool.copyable":               "Copiável",
	"tool.carousel":               "Carrossel",
	"tool.qrcode":                 "Qrcode",
	"tool.field hide":             "Esconder",
	"tool.field display":          "Exibição",
	"tool.table permission":       "Gerar permissões",
	"tool.extra code":             "Código Extra",

	"tool.field display normal":     "Normal",
	"tool.field diplay hide":        "Esconder",
	"tool.field diplay edit hide":   "Editar Esconder",
	"tool.field diplay create hide": "Criar Esconder",

	"tool.generate table model success": "sucesso ao gerar",
	"tool.generate table model fail":    "falha ao gerar",

	"tool.detail display":             "Exibição",
	"tool.detail info":                "Informação detalhada",
	"tool.follow list page":           "Seguir a página da lista",
	"tool.inherit from list page":     "Herdar da página da lista",
	"tool.independent from list page": "Independente da página da lista",

	"generator.query":                 "Query",
	"generator.show edit form page":   "Mostrar formulário de edição",
	"generator.show create form page": "Mostrar formulário de criação",
	"generator.edit":                  "Editar",
	"generator.create":                "Criar",
	"generator.delete":                "Excluir",
	"generator.export":                "Exportar",

	"plugin.plugin":                         "Plugin",
	"plugin.plugin detail":                  "Detalhe do plugin",
	"plugin.introduction":                   "Introdução",
	"plugin.website":                        "Website",
	"plugin.version":                        "Versão",
	"plugin.created at":                     "Criado Em",
	"plugin.updated at":                     "Atualizado Em",
	"plugin.provided by %s":                 "Fornecido por %s",
	"plugin.upgrade":                        "Upgrade",
	"plugin.install":                        "Instalar",
	"plugin.download":                       "Download",
	"plugin.buy":                            "Comprar",
	"plugin.downloading":                    "Baixando",
	"plugin.info":                           "Detalhe",
	"plugin.login":                          "Login",
	"plugin.login to goadmin member system": "Faça login no sistema de membros GoAdmin",
	"plugin.account":                        "Conta",
	"plugin.password":                       "Senha",
	"plugin.learn more":                     "Saber mais",

	"plugin.no account? click %s here %s to register.":    "Sem conta？ Clique %saqui%s para se registrar.",
	"plugin.download fail, wrong name":                    "Falha no download, nome errado",
	"plugin.change to debug mode first":                   "Mude para o modo debug primeiro",
	"plugin.download fail, plugin not exist":              "O download falhou, o plugin não existe",
	"plugin.download fail":                                "Falha no download",
	"plugin.golang develop environment does not exist":    "O ambiente de desenvolvimento de Golang não existe",
	"plugin.download success, restart to install":         "Download bem-sucedido, reinicie para instalar",
	"plugin.restart to install":                           "Reinicie para instalar",
	"plugin.can not connect to the goadmin remote server": "Falha ao conectar o servidor remoto GoAdmin, verifique sua conexão de rede.",

	"admin.basic admin": "Admin básico",
	"admin.a built-in plugins of goadmin which help you to build a crud manager platform quickly.": "Plug-ins integrados ao GoAdmin que ajudam você a construir uma plataforma de gerenciamento crud rapidamente.",
	"admin.official": "Oficial",
}
