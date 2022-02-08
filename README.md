# mutant
## Description
Magneto wants to recruit as many mutants as possible so he can fight the X-Men.
He has hired you to develop a project that detects if a human is a mutant based on his DNA sequence.

## How to run
To run the project locally you have to follow the following steps:
1. First we have to install the dependencies. We have to run `make install`    
2. Move to the `postgresql/` folder and run the command: `docker-compose up --build`
3. Finally, move to the root directory and run `make run`


## Available Endpoints
* __Health Check__
```
curl --request GET \
  --url https://mutant-ms.herokuapp.com/api/mutant-ms/health-check \
  --header 'Content-Type: application/json'
```

* __Is Mutant__
```
curl --request POST \
  --url https://mutant-ms.herokuapp.com/api/mutant-ms/mutant \
  --header 'Content-Type: application/json' \
  --data '{
	"dna": [
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTC "
	]
}'
```

* __Stats__
```
curl --request GET \
  --url https://mutant-ms.herokuapp.com/api/mutant-ms/stats \
  --header 'Content-Type: application/json'
```

## Project link
The link to the project is this [URL](https://mutant-ms.herokuapp.com/)