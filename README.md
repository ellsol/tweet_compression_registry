# Tweet Compression Registry

This service offers (sha3) compression for tweets and also returns the tweet belonging
to a compression

## Dependencies

Prepare

Install

    Goose
    Make (if not present)
    Jet
    Swagger

Run

    make vet
    make migrate
    make jet

## Basics

Run the program

    make run

## Deployment

Build docker container with (update version manually)

    make docker


## Create Docs

    make serve-swagger

## Migration

#### Add New Migration

    goose create NAME_OF_MIGRATION sql
