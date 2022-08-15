## First run on localhost

```shell
make up && sleep 10 && make migrate-local run-local
```
then the API will be available on `localhost:8090`

## Choices
Get all the choices that are usable for the UI.
GET: http://0.0.0.0:8090/choices
Result: application/json
[
  {
    “id": integer [1-5],
    "name": string [12] (rock, paper, scissors, lizard, spock)
  }
]

## Choice
Get a randomly generated choice
GET: http://0.0.0.0:8090/choice
Result: application/json
{
  "id": integer [1-5],
  "name" : string [12] (rock, paper, scissors, lizard, spock)
}

## Play
Play a round against a computer opponent
POST: http://0.0.0.0:8090/play
Data: application/json
{
  “player”: choice_id 
}
Result: application/json
{
  "results": string [12] (win, lose, tie),
  “player”: choice_id,
  “computer”:  choice_id
}

## Scoreboard
 A scoreboard with the 10 most recent results
GET: http://0.0.0.0:8090/scoreboard
Result: application/json
[
	{
		"results": "tie",
		"player": 1,
		"computer": 1,
		"timestamp": "2022-08-15T20:11:15.227571Z"
	},
	{
		"results": "tie",
		"player": 1,
		"computer": 1,
		"timestamp": "2022-08-15T20:11:14.723241Z"
	},

## Reset
The scoreboard to be reset
DELETE: http://0.0.0.0:8090/reset
Result: application/json
{
	"result": true
}