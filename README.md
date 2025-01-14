API AdotaPet
Este é um projeto de uma API REST desenvolvida em Golang para a gestão de adoção de pets. O sistema facilita o cadastro de usuários e pets, permitindo que os usuários se cadastrem, façam login, e busquem, criem, atualizem ou removam pets disponíveis para adoção.

Funcionalidades
Usuários:

-Cadastro de novos usuários com nome, apelido (nick), telefone e senha.

-Login de usuários, com verificação de senha.

-Busca de usuários por nome ou apelido (nick).

-Consulta detalhada de um usuário.

-Atualização e exclusão de usuários.

Pets:

-Cadastro de novos pets para adoção com nome, raça, cor, tamanho, sexo, idade e peso.

-Listagem de todos os pets disponíveis para adoção.

-Consulta detalhada de um pet específico.

-Atualização de informações de um pet.

-Exclusão de um pet.

Segurança:

-Hash de Senha: Utiliza o pacote bcrypt para garantir que as senhas dos usuários sejam armazenadas de forma segura, utilizando um hash seguro.

-Verificação de Senha: A verificação de senha é feita comparando o hash armazenado com a senha fornecida durante o login.

Roteamento e Autenticação
As rotas da API são configuradas e organizadas de maneira modular, utilizando o pacote mux para roteamento.
Cada rota é associada a um controlador, que manipula a lógica das requisições HTTP.
O middleware de logger é aplicado para registrar as requisições feitas à API.
As rotas são organizadas em duas partes principais: Usuários e Pets, com a possibilidade de adicionar mais funcionalidades, como autenticação e autorização, no futuro.

Estrutura de Diretórios
/src/controllers: Contém os manipuladores das requisições HTTP para usuários e pets.
/src/modelos: Define as estruturas de dados (modelos) dos pets e usuários.
/src/repositorios: Responsável pela interação com o banco de dados.
/src/banco: Contém a lógica para conectar ao banco de dados.
/src/middlewares: Implementação de middlewares como logger e CORS.
/src/rotas: Define as rotas da API e a configuração do roteador.
/src/seguranca: Implementação de hash e verificação de senha.

Tecnologias Utilizadas
-Golang: A linguagem de programação principal para construção da API.

-MySQL: Banco de dados utilizado para armazenar as informações de usuários e pets.

-Gorilla Mux: Utilizado para o roteamento das requisições HTTP.

-Godotenv: Para carregar variáveis de ambiente do arquivo .env.

-bcrypt: Para segurança e hash de senhas dos usuários.
