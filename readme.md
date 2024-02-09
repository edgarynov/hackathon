Wiki trieur automatique et créateur arborescence

 
Manuel Utilisateur
 
Comment utiliser le trieur. 

Tout d’abord pour pouvoir utiliser le trieur il vous faudra définir le chemin d’accès des fichiers à trier et ensuite définir le chemin d’arriver des fichiers. 


Fonctionnement trieur
Une fois le chemin d’accès des fichiers à trier est défini en aval et en amont le programme vas commencer par rentrer tous les fichiers dans un bloom filter et leur donnée une identité (le bloom filter actuel est setup pour 500 éléments) mais un risque de faux positif est possible donc j’ai mis la possibilité de faux positif a 0.01% soit une fois sur 10000. Ensuite pour éviter les problèmes liés aux faux positif ou autres bugs qui peuvent subvenir comme un doublon ou la duplication d’un fichier il y a un programme qui vérifie l’identité de tout le fichier a l’aide du hachage MD5 des fichiers et supprime ceux qui ont la même identité qu’un autre sachant que chaque fichier a sa propre identité seul une copie peux avoir la même identité. 

Le trieur peux crée dans le chemin d’accès d’arriver de nouveau sous répertoire et de nouveau dossier en fonction du type et du nom des fichiers 



Documentation Technique 

Le cœur du programme et sa structure est organisé en a ma manière : 
-	Le cœur du programme (sort file.go) est composé de 2 fonction qui vienne entourer et corriger ce que le cœur ne fait pas :
o	Le bloom filter (le filtre qui donne les identité)
o	Le file Identity (le programme vérifiant les doublons et les supprimant à l’aide du hachage MD5) 

Tout est ensuite mit sur le main.go ou ont essayé de faire une interface pour que les chemin d’accès soit + intuitif à rentrer