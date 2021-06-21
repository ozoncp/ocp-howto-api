# OCP Howto API

[![Coverage Actions Status](https://codecov.io/gh/ozoncp/ocp-howto-api/branch/main/graph/badge.svg?token=f8d4ff0d-bedf-4b0d-890c-cb05dc3239e0)](https://codecov.io/gh/ozoncp/ocp-howto-api)

OCP Howto API is an API for Ozon Code Platform that provides access to howto entities.
Howto entity represents pair question-answer in some course.

### Howto 
| Field | Type | Description |
| ------ | ------ | ------ |
| Id | Number | Unique identifier of howto |
| CourseId | Number | Identifier of related course |
| Question | String | Question |
| Answer | String | Answer to the question |

For more information about Ozon Code Platform structure refer to [documentation](https://github.com/ozoncp/docs/).

### Interface

The API contains the following methods:
| Method | Description |
| ------ | ------ |
| AddHowtoV1 | Creates howto in database and returns it's identifier |
| AddHowtosV1 | Creates several howtos and returns identifiers of successfully created entities |
| UpdateHowtoV1 | Updates information about howto in database |
| DecribeHowtoV1 | Returns information about howto  |
| ListHowtosV1 | Returns information about several howtos |
| RemoveHowtoV1 | Removes howto from database |
