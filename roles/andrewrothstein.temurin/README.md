andrewrothstein.temurin
=========

![Build Status](https://github.com/andrewrothstein/ansible-temurin/actions/workflows/build.yml/badge.svg)

Installs Adoptium's OpenJDK distribution [Temurin](https://adoptium.net/releases.html)

Requirements
------------

See [meta/main.yml](meta/main.yml)

Role Variables
--------------

See [defaults/main.yml](defaults/main.yml)

Dependencies
------------

See [meta/main.yml](meta/main.yml)

Example Playbook
----------------

```yml
- hosts: servers
  vars:
    openjdk_app: jre
    openjdk_ver:
      major: 17
      minor: 0
      patch: 1
      b: 12
  roles:
    - andrewrothstein.temurin
```

License
-------

MIT

Author Information
------------------

Andrew Rothstein <andrew.rothstein@gmail.com>
