# API Ticket Management System | Sistema de Gerenciamento de Tickets

<p align="center">
  <img src="static/img/logo.png" alt="API Logo" width="250">
</p>

<div align="center">
  <p>
    <a href="#english">English</a> |
    <a href="#portuguese">Português</a>
  </p>
</div>

---

<a id="english"></a>
# Ticket Management API

<p align="center">
  A complete solution for technical support ticket management
</p>

## About the Project

The Ticket Management API is a complete solution for companies and organizations that need to efficiently manage technical support requests. It offers a comprehensive set of endpoints that enable:

- **User Management**: Create, query, and remove users with ease
- **Ticket Creation and Tracking**: Register new tickets with detailed information
- **Status Updates**: Track ticket progress (pending, in progress, completed)
- **Complete History**: Maintain a detailed record of all updates
- **Advanced Filtering**: Retrieve specific tickets by author or status

## Documentation

For detailed information about all available endpoints, parameters, status codes, and usage examples, see our [complete documentation](DOCUMENTATION.md).

## Main Endpoints

The API offers the following main endpoints:

- **Authentication**: `/auth/enter`, `/auth/register`
- **Users**: `/user/fetch`, `/user/create`, `/user/delete`
- **Tickets**: 
  - Creation: `/ticket/create`
  - Updates: `/ticket/edit`, `/ticket/update`
  - Queries: `/ticket/fetch`, `/ticket/info`
  - Removal: `/ticket/remove`

## Technologies Used

- RESTful API
- JSON for data format
- Support for multiple languages (English and Portuguese)
- Secure authentication

## How to Use

Base URL for all requests:

```
domain/api
```

Example request to create a ticket:

```bash
curl -X POST domain/api/ticket/create \
  -H "Content-Type: application/json" \
  -d '{
    "author_email": "user@example.com",
    "ticket_name": "Printer problem",
    "ticket_label": "Hardware",
    "ticket_equipment": "HP Printer",
    "ticket_explain": "The printer is not connecting to the network"
  }'
```

## License

This project is licensed under the terms of the [MIT License](LICENSE).

---

<a id="portuguese"></a>
# API de Gerenciamento de Tickets

<p align="center">
  <img src="static/img/logo.png" alt="API Logo" width="250">
</p>

<p align="center">
  Uma solução completa para gerenciamento de tickets de suporte técnico
</p>

## Sobre o Projeto

A API de Gerenciamento de Tickets é uma solução completa para empresas e organizações que precisam gerenciar solicitações de suporte técnico de forma eficiente. Ela oferece um conjunto abrangente de endpoints que possibilitam:

- **Gerenciamento de Usuários**: Crie, consulte e remova usuários com facilidade
- **Criação e Rastreamento de Tickets**: Registre novos tickets com informações detalhadas
- **Atualizações de Status**: Acompanhe o progresso dos tickets (pendente, em andamento, concluído)
- **Histórico Completo**: Mantenha um registro detalhado de todas as atualizações
- **Filtragem Avançada**: Recupere tickets específicos por autor ou status

## Documentação

Para informações detalhadas sobre todos os endpoints disponíveis, parâmetros, códigos de status e exemplos de uso, consulte nossa [documentação completa](DOCUMENTATION.md).

## Endpoints Principais

A API oferece os seguintes endpoints principais:

- **Autenticação**: `/auth/enter`, `/auth/register`
- **Usuários**: `/user/fetch`, `/user/create`, `/user/delete`
- **Tickets**: 
  - Criação: `/ticket/create`
  - Atualização: `/ticket/edit`, `/ticket/update`
  - Consulta: `/ticket/fetch`, `/ticket/info`
  - Remoção: `/ticket/remove`

## Tecnologias Utilizadas

- RESTful API
- JSON para formato de dados
- Suporte a múltiplos idiomas (Inglês e Português)
- Autenticação segura

## Como Usar

Base URL para todas as requisições:

```
domain/api
```

Exemplo de requisição para criar um ticket:

```bash
curl -X POST domain/api/ticket/create \
  -H "Content-Type: application/json" \
  -d '{
    "author_email": "usuario@exemplo.com",
    "ticket_name": "Problema com impressora",
    "ticket_label": "Hardware",
    "ticket_equipment": "Impressora HP",
    "ticket_explain": "A impressora não está conectando à rede"
  }'
```

## Licença

Este projeto está licenciado sob os termos da [Licença MIT](LICENSE). 