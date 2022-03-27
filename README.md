# Phone Number Classifier
This service aims to classify phone numbers by Country & State. It also returns the country code & phone number itself.

## API Documentation

<details>
  <summary><b>List Phone Numbers</b></summary>

  </br>

  > **Show all classified phone numbers**

  #### URL
  `/phone`

  #### Method
  `GET`

  ### Query Params
  `country`
  `state`

  #### Success Response
  ```json
  {
      "status": true,
      "result": [
          {
            "country": "Morocco",
            "state": "NOK",
            "countryCode": "+212",
            "PhoneNumber": "6007989253"
          }
      ]
  }
  ```

  #### Try it out
  ```bash
  curl --location --request GET 'localhost:8080/phone'
  curl --location --request GET 'localhost:8080/phone?country=Mozambique&state=ok'
  ```

</details>

### Run In Postman
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4051c5456c226e2bc0ba)

## Developer Guideline

**- Clone Repository**
```
$ git clone https://github.com/ViniciusMartinsS/phone-number-classifier
```

### Useful Commands

**- Run Application On Docker**
```
$ make docker
```

**- Run Application Locally**

<sub>⚠️ Before start, on the `environment` directory, you must create your `app.json` in order to have the environment variables.
On the root of the project, execute the following command:</sup>
```
$ cp environment/app-dist.json environment/app.json
```

<sub> Having the config setup, now you can execute the following commands: <sub>

```
$ make setup
$ make run
```

**- Run Lint Checker**
```
$ make lint
```

**- Run Tests**
```bash
$ make tests
```

**- Generate Test Coverage File**
```bash
$ make coverage
```
