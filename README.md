
# Setup and CRUD golang

## Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Hyper-V or WSL2)
- [Go](https://go.dev/dl/) 1.20+

---

## 1. Start the database
Run TiDB (MySQL-compatible) in a container:

```bash
make db
```
This runs a container named tidb on port 4000 and leaves it running in the background.\
The database will be accessible at 127.0.0.1:4000.

## 2. Apply initial schema
Run:
```bash
make migrate
```
This loads the file db/migrations/001_init.sql into the running database.

## 3. Verify database & tables
```bash
make verify
```

## 4. Configure connection
Create a `.env` file at the root:
```bash
DB_DSN=root:@tcp(127.0.0.1:4000)/tasking?charset=utf8mb4&parseTime=True&loc=Local
```
Add .env to .gitignore if not already ignored.

## 5. Install dependencies
```bash
make deps
```

## 6. Run server
```bash
make run
```

## 7. Stop DB
```bash
make stop-db
```

---
# Questions

## Questions additionnelles

### 1) Quels index sont utilisés dans `/tasks` (GET) et pourquoi ?

- **`tasks.idx_category (category_id)`** : utilisé pour la jointure `tasks.category_id = categories.id` côté table `tasks`.
- **`categories.PRIMARY KEY (id)`** : utilisé pour retrouver rapidement la catégorie correspondante à chaque tâche lors de la jointure.

### 2) Importance de `idx_parent` dans `/tasks/:id/subtasks` ?
- l’index `idx_parent (parent_task_id)` permet à MySQL d’accéder directement aux lignes qui ont ce `parent_task_id`, sans parcourir toute la table `tasks`.
- Comme l’index regroupe les valeurs identiques de `parent_task_id` ensemble et dans l’ordre, on évite un scan complet et on réduit fortement le temps de recherche.


### 3) Comment améliorer la recherche par `status` et `due_date` dans une même requête ?

- Créer un index composite sur `(status, due_date)`.
- L’ordre des colonnes est important :
  - D’abord la colonne filtrée avec une valeur exacte (`status = 'todo'`, `'in_progress'` ou `'done'`).
  - Ensuite la colonne filtrée par `due_date`, qu’il s’agisse d’une date précise ou d’une plage.
- Exemple d’utilisation optimale :
  `WHERE status = 'in_progress' AND due_date BETWEEN '2025-08-01' AND '2025-08-31'`
- Avantage : MySQL lit uniquement les lignes correspondant au `status` choisi, puis parcourt seulement celles correspondant à la date ou à la plage demandée, au lieu de scanner toute la table.


## Culture Informatique

### 1) Qu’est-ce qu’une fonction réentrante ?
Une fonction réentrante est une fonction qui peut être appelée en même temps par plusieurs exécutions (threads, interruptions…) sans provoquer de comportements incorrects.

Pour être réentrante elle doit :
- ne pas modifier de variables globales ou partagées sans protection.
- utiliser uniquement des données locales à chaque appel.

Exemple : une fonction qui calcule une somme à partir de paramètres passés en entrée est réentrante. Une fonction qui modifie un compteur global sans verrouillage ne l’est pas.

### 2) Quelle est la différence entre un thread, un fork et une coroutine ?
- Un fork crée un nouveau processus indépendant avec sa propre mémoire (lourd).
- Un thread s’exécute en parallèle dans le même processus et partage la mémoire.
- Une coroutine est une exécution légère dans le même processus qui s’alterne avec d’autres coroutines sans faire de parallélisme matériel comme un thread.

### 3) Qu’est-ce que HELM ?
HELM est un gestionnaire de packages pour Kubernetes.
- Permet de déployer, configurer et mettre à jour des applis dans un cluster Kubernetes.
- HELM gère aussi les versions et facilite le rollback de version.

On peut le voir comme le npm de Kubernetes.\
Je n'ai personnellement encore jamais travaillé avec Kubernetes.

### 4) Que pensez-vous des design patterns ? Avez-vous déjà utilisé un ou plusieurs patron(s) de conception ?

Je pense que ce sont des solutions éprouvées pour résoudre des problèmes récurrents dans la conception logicielle.Bien utilisés, ils rendent le code plus clair, cohérent et maintenable.

Exemples :

- Singleton : garantir une seule instance d’un logger global, pour éviter les doublons et assurer une configuration uniforme.

- Observer : Dans un chat, tous les clients connectés (navigateur web, appli mobile, appli desktop) reçoivent instantanément le nouveau message sans qu’on ait à les appeler un par un dans le code.

- Factory : créer la bonne connexion à une base de données (MySQL, PostgreSQL, etc...) sans que l’appelant ait besoin de connaître les détails techniques.

J’ai déjà manipulé le principe du Singleton dans des contextes réels, et je suis à l’aise avec l’idée d’adopter d’autres patterns selon les besoins.

### 5) Qu-est-ce que le NewSQL ?
C'est une nouvelle génération de bases de données relationnelles conçues pour offrir la scalabilité des bases NoSQL tout en gardant les avantagezs du SQL (langage SQL, transactions ACID).

Exemples : TiDB, CockroachDB, Google Spanner.