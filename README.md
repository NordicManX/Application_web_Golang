# 🏬 Loja Web com Go

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql)
![Bootstrap](https://img.shields.io/badge/Bootstrap-4.3-7952B3?style=for-the-badge&logo=bootstrap)

Uma aplicação web simples para gerenciamento de produtos, construída com Go e PostgreSQL. Este projeto serve como um exemplo prático de uma aplicação CRUD (Create, Read, Update, Delete) com renderização de templates no lado do servidor.

<br>

## 📋 Tabela de Conteúdos
- [📖 Sobre o Projeto](#-sobre-o-projeto)
- [✨ Features](#-features)
- [🛠️ Tecnologias Utilizadas](#️-tecnologias-utilizadas)
- [🚀 Começando](#-começando)
  - [Pré-requisitos](#-pré-requisitos)
  - [Instalação e Execução](#️-instalação-e-execução)
- [🌐 Rotas da Aplicação](#-rotas-da-aplicação)
- [📂 Estrutura de Arquivos](#-estrutura-de-arquivos)
- [📄 Licença](#-licença)

<br>

## 📖 Sobre o Projeto
Este projeto implementa uma loja virtual básica onde é possível listar, criar, editar e excluir produtos. A arquitetura separa as responsabilidades em pacotes, seguindo as convenções da comunidade Go para projetos web.

## ✨ Features
-   [x] **Listar Produtos:** Visualização de todos os itens do inventário.
-   [x] **Criar Produtos:** Adicionar novos itens através de um formulário web.
-   [x] **Editar Produtos:** Atualizar informações de produtos existentes.
-   [x] **Excluir Produtos:** Remover itens do inventário.

## 🛠️ Tecnologias Utilizadas
- **Backend:** Go (Golang)
- **Banco de Dados:** PostgreSQL
- **Frontend:** HTML5 com Go Templates
- **Estilização:** Bootstrap 4

## 🚀 Começando
Siga estas instruções para obter uma cópia do projeto e executá-la em sua máquina local.

### 📋 Pré-requisitos
Certifique-se de ter os seguintes softwares instalados:
- [Go](https.go.dev/doc/install) (versão 1.18 ou superior)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

### ⚙️ Instalação e Execução

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)
    cd seu-repositorio
    ```

2.  **Configure o Banco de Dados:**
    Acesse seu PostgreSQL e crie a tabela `produtos` executando o script abaixo:
    ```sql
    CREATE TABLE produtos (
        id SERIAL PRIMARY KEY,
        nome VARCHAR(255) NOT NULL,
        descricao TEXT,
        preco NUMERIC(10, 2) NOT NULL,
        quantidade INTEGER NOT NULL
    );
    ```

3.  **Configure as Variáveis de Ambiente:**
    Crie um arquivo `.env` na raiz do projeto (você pode copiar o `.env.example` se houver um) e preencha com suas credenciais do banco de dados:
    ```ini
    # .env
    DB_USER=seu_usuario_postgres
    DB_PASSWORD=sua_senha_segura
    DB_NAME=o_nome_do_banco
    DB_HOST=localhost
    DB_PORT=5433
    ```
    *(Lembre-se de adicionar `.env` ao seu arquivo `.gitignore`)*

4.  **Instale as dependências do projeto:**
    ```bash
    go mod tidy
    ```

5.  **Execute a aplicação:**
    ```bash
    go run main.go
    ```
    🎉 A aplicação estará disponível em `http://localhost:8000`.

## 🌐 Rotas da Aplicação

| Método | Rota         | Descrição                                         |
| :------- | :----------- | :-------------------------------------------------- |
| `GET`    | `/`          | Exibe a página inicial com a lista de produtos.   |
| `GET`    | `/new`       | Exibe o formulário para criar um novo produto.    |
| `POST`   | `/insert`    | Processa os dados do novo produto e o salva.      |
| `GET`    | `/edit`      | Exibe o formulário de edição para um produto.     |
| `POST`   | `/update`    | Processa a atualização dos dados do produto.      |
| `GET`    | `/delete`    | Deleta um produto específico.                     |


## 📂 Estrutura de Arquivos

.
├── controllers/    # Lógica de requisição e resposta HTTP
├── db/             # Conexão com o banco de dados
├── models/         # Estruturas de dados e acesso ao banco
├── routes/         # Definição das rotas da API
├── templates/      # Arquivos HTML
└── main.go         # Ponto de entrada da aplicação


## 📄 Licença

Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.
