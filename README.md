# Phone Number Classifier
This service aims to classify phone numbers by Country & State. It also returns the country code & phone number itself.

## API Documentation

<details>
  <summary><b>List Tasks</b></summary>

  </br>

  > **Show all task of a technician**

  #### URL
  `/phone`

  #### Method
  `GET`

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

  #### Error Response
  ```json
  {
    "status": false,
    "message": "Internal Server Error Occurred. Please, get in touch with the support team."
  }
  ```

  #### Try it out
  ```bash
  curl --location --request GET 'localhost:8080/phone'
  ```

</details>

### Run In Postman
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/955776fb0c44d8d9235f)

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
