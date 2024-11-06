# s05-Player


## Le challenge pour ce soir :

### Implémenter une structure Player contenant les données suivantes : nom, pseudo, age (int), health(int), mana (int)
ok 

### Déclarer et initialiser une map players qui index avec le nom d'un player des objects de type *Player
ok 

### Implémenter pour la structure Player les fonctions suivantes :
ok
#### save() qui permet de stocker dans un fichier nommé nom.yml toutes les données d'un joueur
ok
#### del() qui permet de supprimer un player (de la map et qui supprimer le fichier .yml avec son tests unitaire associé
ok 

#### display() qui retourne une string décrivant un joueur avec son test unitaire associé
ok 

#### Une fonction playerLoad(name string) qui retourne un Player et qui va :
ok
#### charger un player depuis la map s'il existe dedans
ok
#### créer un player s'il n'existe pas en demandant une saisie utilisateur (seulement le pseudo) sur l'entrée standard (stdin) puis l'ajouter dans la map
ok
#### charger le player en question s'il existe depuis le fichier .yml et l'ajouter a la map
ok 

### Un test unitaire qui load un player unexistant, le créé et le détruit ensuite.

### Note : en utilisant les tests unitaire uniquement pour vos tests vous pouvez éviter d'implémenter la fonction main() qu'on met généralement dans le main.


### go et donc vous contenter pour ce challenge de deux fichiers player.go et player_test.go