# E-Commerce

# Index

<!--ts-->
   * [Solução](#solução)
   * [Detalhadamente](#detalhadamente)
   * [Explicação de algumas decissões](#explicação-de-algumas-decissões)
   * [Pontos de melhoria](#pontos-de-melhoria)
   * [Executando a aplicação](#executando-a-aplição)
   * [Bônus](#bônus)
<!--te-->

---

Solução
=======
Consistente em expor uma entry-port para realização do checkout de um E-Commerce através de um endpoint HTTP.

---

Detalhadamente
==============
O design arquitetural da solução é dividida em dois grandes módulos, 
se baseando em conceitos de ports and adapters e monolito modular:

## Core
É a parte que representa as regras de negócio, composto por módulos de domínio que compõem e formam a API da solução,
upstream praticamente para todo projeto.

A granularidade optada para os módulos é a nível de sub-dominios, composto por:

### Domain
Onde estão os componentes que compõem as regras de negócio módulo.

### Application
Onde estão as funcionalidades de cada módulo estão, encapsulando, orquestrando todos componentes, 
e complexidades necessárias, expondo as suas funcionalidades através de api's bem definidas.

Downstream em relação ao Domain de seu módulo, mas upstream para api de outros módulos, entry-ports como endpoints HTTP 
ou outra extremidade que precise interfacear com uma de suas funcionalidades.

## Platform
É a parte base da infraestrutura que provê recursos que juntos ao Core completa a solução.

### Configuration
Definição e carregamento da configuração da aplicação baseada em DOTENV.

### Dependency Injection
Container de dependências que contém configurado por módulo todo o grafo de dependências da aplicação.

### Infrastructure
Componentes que dão base e suporte para a solução, como os adapters/out-ports.

### HTTP
Camada de transporte baseada no protocolo HTTP, onde expõe as funcionalidades através dos entry-ports, 
formando a Web API.

### Componente Run
Componente que encapsula e orquestra todo processo para execução da aplicação através da plataforma.

É a API chamada ao executar a solução.

---

Explicação de algumas decissões
===============================

## Integração com um serviço GRPC
Foi o meu primeiro contato real utilizando GRPC, logo acredito que é importante um aprofundamento no tema, 
levando ao ponto crucial, que é referente a forma como tratei a resiliência, provavelmente está longe de ser o melhor, 
mas para o escopo do desafio atendeu. 

Para mim o correto como próximo passo seria entender melhor o que o próprio GRPC já oferece como comportamento 
resiliente, caso não fosse o ideal ou suficiente, buscaria implementar técnicas como Retry, Circuit-Breaker e até mesmo 
Cache, caso o valor de desconto de cada produto tivesse uma constância que permitisse, como por exemplo ser o 
mesmo sempre que solicitar um desconto para um mesmo produto.

## Aplicação de desconto
Como o valor monetário está sendo trabalhado na unidade de centavos, sendo um valor inteiro, 
ao receber um desconto que pode ser decimal, ao calcular o desconto, 
eu optei por eliminar a parte decimal, pois são milésimos de real.

## Promoções
Sem muito contexto no domínio de E-commerce, mas utilizando a experiência de cliente de alguns, imaginei que poderiam 
haver diversas promoções ativas ao mesmo tempo, logo existindo um fluxo dinâmico como um pipeline de promoções por onde 
o carrinho de compras passaria e seria modificado através das promoções que fizessem sentido, bastando configurar esse
pipeline de promoções e suas regras de aplicação.

Foi a forma que encontrei em lidar com a dinamicidade de diferentes promoções com diferentes regras e manter essa 
complexidade em um fluxo coeso.

---

Pontos de melhoria
==================

## Promoções
Analisaria e estudaria a possibilidade de evoluir:

- Os componentes envolvidos para uma estrutura de dados mais adequada e execução de uma forma mais eficiente do que 
apenas iterar uma coleção, por exemplo, evitar tentar aplicar promoções não aplicáveis em certos contextos, podendo 
Promotions evoluir para um próprio subdomain.
- Tornar realmente configurável.
- Utilizar feature toggle para que pudéssemos ativar e desativar promoções de forma flexível, sem que precisamos 
intervir no runtime das instâncias.

## Black Friday
Trabalhei com data por ser um requisito do desafio, mas optaria de utilizar feature toggle, onde uma vez a promoção 
da Black Friday estivesse configurada na aplicação, bastaria eu ativar e desativar quando fosse necessário, sem que 
fosse necessária intervir na aplicação antes ou depois da Black Friday.

## Testes
- Definir uma estrutura de dados base para os TestCases de cada componente.
- Encapsular certos dados e comportamentos duplicados níveis acima para compartilhar entre os testes dependentes.
- Organizar melhor os diferentes tipos de testes.
- Algum outro caso de teste ou tipo de teste que possa ter faltado.

## Logs
- Analisar a possibilidade de haver redundância ou a falta de informação a ser logada e melhorar.
- Analisar a necessidade e como estruturar melhor as informações logadas.

## Geral
- Encapsular alguma informação que pode ter ficado duplicada ou pode ser abstraída.
- Encapsular algum comportamento que pode ter ficado duplicado ou pode ser abstraído.
- Refatorar o nome de alguma parte de toda implementação para uma melhor legibilidade, representação do conceito ou até 
mesmo erro no inglês que acabei percebendo depois.
- Alguma técnica, lógica, algoritmo melhor que poderia ser utilizado e até algum recurso da linguagem.
- Manter evolutiva.

---

Executando a aplicação
======================

1- Montar o arquivo de configuração .env na raiz do projeto a partir do base.env, caso ainda não tenha:

    - BLACK_FRIDAY_DAY:

      Variável para configurar o dia em que a promoção da Black Friday deve ser aplicada.
      
      Baseado no layout RFC3339 ex: 
      2021-09-28T00:00:00Z

    - GRPC_SERVER_ADDRESS

      Variável para configurar endereço onde encontrar o discount-service.
      Ex: localhost:50051

    - HTTP_SERVER_PORT

      Variável para configurar a porta que o endpoint HTTP estará exposta.
      Ex: :5000
      
1.2- As configurações definidas no .env e Dockerfile, devem estar alinhadas.

2- Rode a solução executando o docker-compose na raiz do projeto:

    sudo docker-compose up --build

Após ter construido a imagem do E-Commerce a primeira vez, não é necessario mais utilizar o paramatro --build.

3- É esperado os seguintes logs ao rodar os containers sem o modo Detached:

<a href="https://ibb.co/rF9nVBt"><img src="https://i.ibb.co/pbqCSHP/ecommerce.png" alt="ecommerce" border="0"></a>

## Web Api

<details><summary>Checkout</summary>
<p>

```
Endpoint: /checkout
HTTP verb: POST
Payload:
{
    "products": [
        {
            "id": 1,
            "quantity": 1 // Quantidade a ser comprada do produto
        }
    ]
}
```
</p>
</details>

### Comportamento

<details><summary>Caso de sucesso.</summary>
<p>

```
Code: 200
Descripion: Sucesso
Response:
{
  "total_amount": 15157,
  "total_amount_with_discount": 15150,
  "total_discount": 7,
  "products": [
    {
      "id": 1,
      "quantity": 1,
      "unit_amount": 15157,
      "total_amount": 15157,
      "discount": 7,
      "is_gift": false
    }
  ]
} 
```

</p>
</details>

<details><summary>Caso de sucesso na Black Friday.</summary>
<p>

```
Request:
{
    "products": [
        {
            "id": 1,
            "quantity": 1 // Quantidade a ser comprada do produto
        },
        {
            "id": 6,
            "quantity": 1 // Quantidade a ser comprada do produto
        },
    ]
}

Response
Code: 200
Descripion: Sucesso
Payload:
{
  "total_amount": 16057,
  "total_amount_with_discount": 16050,
  "total_discount": 7,
  "products": [
    {
      "id": 1,
      "quantity": 1,
      "unit_amount": 15157,
      "total_amount": 15157,
      "discount": 7,
      "is_gift": false
    },
    {
      "id": 6,
      "quantity": 2,
      "unit_amount": 900,
      "total_amount": 900,
      "discount": 0,
      "is_gift": false
    }
  ]
} 
```

</p>
</details>

<details><summary>Não selecionar nenhum produto.</summary>
<p>

```
Request:
{
  "products": []
}
 
Response
Code: 200
Descripion: Sucesso
Payload:
{
  "total_amount": 0,
  "total_amount_with_discount": 0,
  "total_discount": 0,
  "products": []
} 
```

</p>
</details>

<details><summary>Ignora produtos repetidos.</summary>
<p>

```
Request:
{
  "products": [
    {
      "id": 1,
      "quantity": 1
    },
    {
      "id": 1,
      "quantity": 1
    }
 ]
}
 
Response
Code: 200
Descripion: Sucesso
Payload:
{
  "total_amount": 15157,
  "total_amount_with_discount": 15150,
  "total_discount": 7,
  "products": [
    {
      "id": 1,
      "quantity": 1,
      "unit_amount": 15157,
      "total_amount": 15157,
      "discount": 7,
      "is_gift": false
    }
  ]
} 
```

</p>
</details>

<details><summary>Paylaod inválido.</summary>
<p>

```
Request:
{
  {
    "id": 1,
    "quantity": 1
  },
  {
    "id": 2,
    "quantity": 1
  }
 ]
} 

Response
Code: 400
Descripion: Bad request
Payload: -
```

</p>
</details>

<details><summary>Produtos selecionados com valor invalido, como informar zero, negativo e decimal são ignorados.</summary>
<p>

```
Request:
{
  "products": [
    {
      "id": -1,
      "quantity": 1.10
    },
    {
      "id": 0,
      "quantity": 0
    },
    {
      "id": 1,
      "quantity": 1
    }
 ]
} 

Response
Code: 200
Descripion: Success
Payload:
{
  "total_amount": 15157,
  "total_amount_with_discount": 15150,
  "total_discount": 7,
  "products": [
    {
      "id": 1,
      "quantity": 1,
      "unit_amount": 15157,
      "total_amount": 15157,
      "discount": 7,
      "is_gift": false
    }
  ]
} 
```

</p>
</details>

---

Bônus
=====
Projeto pessoal que tenho trabalhado atualmente.

## Core Pay
Um sistema central(core) para soluções financeiras(vulgo fintechs) comuns ou disruptivas, 
servindo como a principal interface para operações financeiras base.

Mudando o foco de produto para plataforma, o projeto e-Money foi reformulado, dando lugar para o Core Pay.

- [Artigo](https://dev.to/gmarcial/core-pay-solucao-5af5)
- [Repositório](https://github.com/corelab1/corepay)
- [Docs](https://app.gitbook.com/@guilherme-marcial-felipe/s/core-pay/)
