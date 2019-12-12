# OKAPI V2 - Go validation project

## Clean Architecture to validate go

quelques use case du planning :

### Use cases

Planning : 
- x Creer le Planning d'un agent : Liste de tache
- x Creer des taches : 
  - x Nom de tache
  - x Date debut
  - x Date fin
  - x Type (String)
- x Ajouter une tache a un planning : 
  - validation du non overlap des taches sur le temps
  (overlap = vérification que deux tâches ne se chevauchent pas dans le PLanning)
- Ajouter une absence :
  - Type de tache particulier

### presentation 

1 mode API pour gerer ca
1 mode console

### persitence

1 base de donnee sqlite
1 base in memory





