#! /bin/bash

id_number=$(curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Nathan\"}}){id}}"}')
echo $id_number