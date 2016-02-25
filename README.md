# Dev Shop

![Imgur](http://i.imgur.com/8NPz67T.png)

Dev shop é uma aplicação de teste e aprendizagem.

Utiliza uma API web em [Go](http://golang.org), persistencia com [LevelDB](https://github.com/google/leveldb), e front-end com [VUEJS](vuejs.org) e [Material Design Lite](http://getmdl.io).

[Veja ao vivo](http://104.236.203.153).

## Desenvolvimento

O ambiente de desenvolvimento é configurado no [docker](https://www.docker.com/) usando [docker-compose](https://docs.docoker.com/compose), então idealmente você vai precisar de ambos.

Para os comandos no docker existem tasks no Makefile.

    $ make deps # instala dependencias
    $ make run # inicia o servidor
    $ make run-dev # inicia o servidor e reinicia a cada 10s
    $ make tests  # roda os testes
    $ make build # constroí imagem docker para produção
    $ make push # envia imagem docker para dockerhub

Também pode ser feito normalmente usando o GO instalado localmente.

    $ go tests
    $ ./run.sh

Por padrão o sistema faz chamadas a API do github sem token, logo caí rápido no rate limiting. Para usar um token basta criar um arquivo .token na raiz do projeto cujo o conteúdo é o seu API TOKEN.

## License 

MIT

