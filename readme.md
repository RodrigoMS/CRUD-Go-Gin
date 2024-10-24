# Aplicação CRUD com Golang e PostgreSQL
    
Esta aplicação é um exemplo de implementação de um CRUD (Create, Read, Update, Delete) utilizando Golang, PostgreSQL e Docker. A aplicação segue o modelo de arquitetura MVC (Model-View-Controller).

## Funcionalidades

- Listar todos os items
- Criar um novo item
- Obter um item por ID
- Atualizar um item
- Deletar um item por ID

## Arquitetura e Conceitos

- **Golang**: Linguagem de programação utilizada para implementar a aplicação.
- **PostgreSQL**: Sistema de banco de dados relacional utilizado para armazenar os dados.
- **Gin Gonic**: Framework web para Golang utilizado para criar rotas e lidar com requisições HTTP.
- **Docker**: Ferramenta utilizada para criar contêineres de desenvolvimento e produção.
- **Docker Compose**: Ferramenta para definir e gerenciar multi-contêineres Docker.

## Pré-requisitos

- Docker e Docker Compose instalados

## Comandos para Iniciar a Aplicação

### Clonar o repositório
```bash
git clone https://github.com/RodrigoMS/CRUD-Go-Gin.git
cd CRUD-Go-Gin
```

### Construir e rodar a aplicação com Docker Compose

```bash
docker-compose up --build
```    

### Parar e remover os containers

```bash
docker-compose down
```

### Endpoints da API

*   **Listar todos os items**
    
    http
    
    ```bash
    
        GET /items
       ``` 
    
*   **Criar um novo item**
    
    http
    
    ```bash
        POST /items
        Content-Type: application/json
        
        {
            "name": "Item1",
            "price": 100
        }
    ```
    
*   **Obter um item por ID**
    
    http
    
    ```bash
        GET /items/:id
    ``` 
    
*   **Atualizar um item**
    
    http
    
    ```bash
        PUT /items/:id
        Content-Type: application/json
        
        {
            "name": "ItemUpdated",
            "price": 150
        }
    ```
    
*   **Deletar um item por ID**
    
    http
    
    ```bash
        DELETE /items/:id
    ```
    

Bibliotecas Utilizadas
----------------------

*   [Golang](https://golang.org/)
    
*   [Gin Gonic](https://github.com/gin-gonic/gin)
    
*   [PostgreSQL](https://www.postgresql.org/)
    
*   [Docker](https://www.docker.com/)
    
*   [Docker Compose](https://docs.docker.com/compose/)


Extensões VSCode
----------------
*   Rest Client - by Huachao Mao

*   Go - by Go Team at Google
