go-gettokenbyserviceaccount
=====

[![Build Status](https://travis-ci.org/tanaikech/go-gettokenbyserviceaccount.svg?branch=master)](https://travis-ci.org/tanaikech/go-gettokenbyserviceaccount)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENCE)

<a name="TOP"></a>
# Overview
This is a Golang library to retrieve access token from [Service Account of Google](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) without using [Google's OAuth2 package](https://github.com/golang/oauth2).

# Install
You can install this using ``go get`` as follows.

~~~bash
$ go get -u github.com/tanaikech/go-gettokenbyserviceaccount
~~~

# Usage

~~~
res, err := gettokenbyserviceaccount.Do(privateKey, clientEmail, scopes)
~~~

- ``privateKey``, ``clientEmail`` and ``scopes`` are string values.

You can obtain the access token like below.

~~~
{
  "access_token": "#####",
  "expires_in": 3600,
  "token_type": "Bearer",
  "start_time": 1234567890,
  "end_time": 1234567890
}
~~~

You can also retrieve this result using Google's OAuth2 package. I created this library to study the JWT process.

-----

<a name="Licence"></a>
# Licence
[MIT](LICENCE)

<a name="Author"></a>
# Author
[Tanaike](https://tanaikech.github.io/about/)

If you have any questions and commissions for me, feel free to tell me.

<a name="Update_History"></a>
# Update History
* v1.0.0 (December 11, 2018)

    1. Initial release.


[TOP](#TOP)
