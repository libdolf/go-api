# API Simples em Go - README

---

## Descrição

Esta é uma API simples em Go que implementa dois endpoints REST básicos para manipulação de usuários. Os endpoints são `/users/` para operações CRUD e `GET /users/{id}` para recuperar informações de um usuário específico.

---

## Tecnologias Utilizadas

- **Go:** A linguagem de programação principal para o desenvolvimento.
- **HTTP:** Implementação nativa em Go para criar um servidor HTTP básico.
- **JSON:** Formato de intercâmbio de dados utilizado para representar objetos de usuário.

---

## Estrutura do Projeto

- **`main.go`:** Arquivo principal que inicia o servidor e lida com as rotas.
- **`user.go`:** Define a estrutura de dados `User` e funções relacionadas a usuários.
- **`handlers.go`:** Implementa os manipuladores para os endpoints `/users/` e `GET /users/{id}`.
- **`utils.go`:** Funções auxiliares e utilitárias.

---

## Endpoints

1. **`GET /users/`**
   - **Descrição:** Retorna a lista de todos os usuários.
   - **Exemplo de Requisição:**
     ```bash
     curl -X GET http://localhost:8080/users/
     ```
   - **Exemplo de Resposta:**
     ```json
     [
         {"ID": "1", "Name": "Alice"},
         {"ID": "2", "Name": "Bob"}
     ]
     ```

2. **`GET /users/{id}`**
   - **Descrição:** Retorna as informações de um usuário específico com o ID fornecido.
   - **Exemplo de Requisição:**
     ```bash
     curl -X GET http://localhost:8080/users/1
     ```
   - **Exemplo de Resposta:**
     ```json
     {"ID": "1", "Name": "Alice"}
     ```

     ## Endpoints

1. **`GET /users/`**
   - **Descrição:** Retorna a lista de todos os usuários.
   - **Exemplo de Requisição:**
     ```bash
     curl -X GET http://localhost:8080/users/
     ```
   - **Exemplo de Resposta:**
     ```json
     [
         {"ID": "1", "Name": "Alice"},
         {"ID": "2", "Name": "Bob"}
     ]
     ```

2. **`GET /users/{id}`**
   - **Descrição:** Retorna as informações de um usuário específico com o ID fornecido.
   - **Exemplo de Requisição:**
     ```bash
     curl -X GET http://localhost:8080/users/1
     ```
   - **Exemplo de Resposta:**
     ```json
     {"ID": "1", "Name": "Alice"}
     ```

3. **`POST /users/`**
   - **Descrição:** Adiciona um novo usuário.
   - **Exemplo de Requisição:**
     ```bash
     curl -X POST -H "Content-Type: application/json" -d '{"ID": "3", "Name": "Charlie"}' http://localhost:8080/users/
     ```
   - **Exemplo de Resposta:**
     ```json
     {"ID": "3", "Name": "Charlie"}
     ```

---

## Executando a Aplicação

1. **Instalação do Go:**
   Certifique-se de ter o Go instalado em sua máquina.

2. **Clonando o Repositório:**
   ```bash
   git clone https://github.com/seu-usuario/sua-api-go.git
   cd sua-api-go
3. **Rodando o projeto**
 ```bash
go run main.go
