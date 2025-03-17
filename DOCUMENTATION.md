# API de Gerenciamento de Tickets

Esta API permite gerenciar tickets de suporte técnico, incluindo autenticação de usuários, criação, edição, consulta e remoção de tickets.

## Base URL
```
domain/api
```

## Autenticação

### Login
**Endpoint:** `POST /auth/enter`

**Descrição:** Autentica um usuário no sistema, verificando suas credenciais e permitindo acesso às funcionalidades protegidas da API.

**Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "User has been loged in",
    "status": true
}
```
- Email não registrado (400):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```
- Senha incorreta (400):
```json
{
    "message": "Password is incorrect",
    "status": false
}
```

### Registro
**Endpoint:** `POST /auth/register`

**Descrição:** Registra um novo usuário no sistema, permitindo que ele crie uma conta para acessar a plataforma.

**Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "User has been registered",
    "status": true
}
```
- Email já existe (400):
```json
{
    "message": "Email already exists",
    "status": false
}
```
- Senha inválida (400):
```json
{
    "message": "Password is invalid",
    "status": false
}
```

## Usuários

### Obter Informações do Usuário
**Endpoint:** `GET /user/fetch`

**Descrição:** Recupera informações de um usuário específico ou lista todos os usuários do sistema.

**Query Parameters:**
- `email`: Email do usuário (opcional, exemplo: `email=johndoe@example.com`)

**Respostas:**
- Sucesso (200) - Usuário específico:
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "********",
    "created_at": "2023-01-01T12:00:00Z",
    "status": true
}
```
- Sucesso (200) - Lista de usuários:
```json
{
    "users": [
        {
            "username": "John Doe",
            "email": "johndoe@example.com"
        },
        {
            "username": "Jane Smith",
            "email": "janesmith@example.com"
        }
    ],
    "status": true
}
```
- Usuário não encontrado (404):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```

### Criar Usuário
**Endpoint:** `POST /user/create`

**Descrição:** Cria um novo usuário no sistema.

**Request Body:**
```json
{
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "Password123"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "User has been created",
    "status": true
}
```
- Email já existe (400):
```json
{
    "message": "Email already exists",
    "status": false
}
```
- Dados inválidos (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Remover Usuário
**Endpoint:** `POST /user/delete`

**Descrição:** Remove um usuário do sistema.

**Request Body:**
```json
{
    "email": "johndoe@example.com"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "User has been deleted",
    "status": true
}
```
- Usuário não encontrado (404):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```

## Tickets

### Criar Ticket
**Endpoint:** `POST /ticket/create`

**Descrição:** Cria um novo ticket de suporte técnico.

**Request Body:**
```json
{
    "author_email": "johndoe@example.com",
    "ticket_name": "Problema com impressora",
    "ticket_label": "Hardware",
    "ticket_equipment": "Impressora HP LaserJet",
    "ticket_explain": "A impressora não está conectando à rede"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "Ticket has been created",
    "status": true
}
```
- Dados inválidos (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```
- Autor não encontrado (400):
```json
{
    "message": "Email aren't registered",
    "status": false
}
```

### Atualizar Status do Ticket
**Endpoint:** `POST /ticket/edit`

**Descrição:** Atualiza o status de um ticket existente.

**Request Body:**
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
    "ticket_status": "doing"
}
```

**Observações:**
- Os valores válidos para `ticket_status` são: `pending`, `doing`, `conclued`

**Respostas:**
- Sucesso (200):
```json
{
    "message": "Ticket has been edited",
    "status": true
}
```
- Ticket não encontrado (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
- Status inválido (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Atualizar Histórico do Ticket
**Endpoint:** `POST /ticket/update`

**Descrição:** Adiciona uma nova entrada no histórico de um ticket existente.

**Request Body:**
```json
{
    "ticket_id": "ticket_028492wsd88178",
    "ticket_return": "Comprando roteadores"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "Update has been setting up",
    "status": true
}
```
- Ticket não encontrado (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```

**Observações:**
- O histórico atualizado pode ser consultado através do endpoint `/ticket/info`
- Cada atualização é registrada com a data e hora em que foi realizada
- O histórico é exibido em ordem cronológica

### Listar Tickets
**Endpoint:** `GET /ticket/fetch`

**Descrição:** Lista tickets de um autor específico ou todos os tickets do sistema.

**Query Parameters:**
- `author`: Email do autor (opcional, exemplo: `author=johndoe@example.com`)
- `status`: Status dos tickets (opcional, exemplo: `status=pending`)

**Observações:**
- Os valores válidos para `status` são: `pending`, `doing`, `conclued`
- Se o parâmetro `status` não for fornecido, serão retornados tickets de todos os status
- Se o parâmetro `author` não for fornecido, serão retornados tickets de todos os autores

**Exemplos de Uso:**
- Listar todos os tickets: `/ticket/fetch`
- Listar tickets pendentes: `/ticket/fetch?status=pending`
- Listar tickets de um autor: `/ticket/fetch?author=johndoe@example.com`
- Listar tickets pendentes de um autor: `/ticket/fetch?author=johndoe@example.com&status=pending`

**Respostas:**
- Sucesso (200):
```json
{
    "tickets": [
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
            "ticket_name": "Problema com impressora",
            "ticket_status": "pending"
        },
        {
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174001",
            "ticket_name": "Problema com monitor",
            "ticket_status": "doing"
        }
    ],
    "status": true
}
```
- Nenhum ticket encontrado (404):
```json
{
    "message": "Author don't have tickets",
    "status": false
}
```

### Obter Informações do Ticket
**Endpoint:** `GET /ticket/info`

**Descrição:** Recupera informações detalhadas de um ticket específico, incluindo seu histórico completo.

**Query Parameters:**
- `ticket_id`: ID do ticket (obrigatório)

**Respostas:**
- Sucesso (200):
```json
{
    "ticket_id": "ticket_028492wsd88178",
    "ticket_name": "Problema com roteadores",
    "ticket_status": "doing",
    "ticket_explain": "Necessário comprar novos roteadores",
    "history": [
        {
            "ticket_return": "Purchasing routers",
            "ticket_date": "2023-07-15T14:30:45Z"
        },
        {
            "ticket_return": "Configuring notebook",
            "ticket_date": "2023-07-16T09:15:22Z"
        },
        {
            "ticket_return": "Clean screen",
            "ticket_date": "2023-07-17T11:45:30Z"
        }
    ],
    "status": true
}
```
- Ticket não encontrado (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
- ID não fornecido (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

### Remover Ticket
**Endpoint:** `POST /ticket/remove`

**Descrição:** Remove um ticket do sistema. Apenas o autor do ticket pode removê-lo.

**Request Body:**
```json
{
    "author_email": "johndoe@example.com",
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000"
}
```

**Respostas:**
- Sucesso (200):
```json
{
    "message": "Ticket has been removed",
    "status": true
}
```
- Ticket não encontrado ou autor incorreto (404):
```json
{
    "message": "Ticket hasn't found",
    "status": false
}
```
- Dados inválidos (400):
```json
{
    "message": "Dados inválidos",
    "status": false
}
```

## Códigos de Status

A API utiliza os seguintes códigos de status HTTP:

- `200 OK`: A requisição foi bem-sucedida
- `400 Bad Request`: A requisição contém dados inválidos ou está mal formatada
- `404 Not Found`: O recurso solicitado não foi encontrado
- `500 Internal Server Error`: Ocorreu um erro interno no servidor

## Modelos de Dados

### Usuário
```json
{
    "id": 1,
    "username": "John Doe",
    "email": "johndoe@example.com",
    "password": "********",
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z"
}
```

### Ticket
```json
{
    "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
    "author_id": 1,
    "ticket_name": "Problema com impressora",
    "ticket_label": "Hardware",
    "ticket_equipment": "Impressora HP LaserJet",
    "ticket_explain": "A impressora não está conectando à rede",
    "ticket_status": "pending",
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z",
    "history": [
        {
            "id": "123e4567-e89b-12d3-a456-426614174000",
            "ticket_id": "ticket_123e4567-e89b-12d3-a456-426614174000",
            "ticket_return": "Verificação inicial realizada",
            "ticket_date": "2023-01-01T12:00:00Z"
        }
    ]
}
```

### Status de Ticket
Os tickets podem ter os seguintes status:
- `pending`: Ticket pendente, aguardando atendimento
- `doing`: Ticket em andamento, sendo atendido
- `conclued`: Ticket concluído, atendimento finalizado 