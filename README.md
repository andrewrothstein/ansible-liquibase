# Liquibase Ansible Role

![Build Status](https://github.com/liquibase/liquibase/actions/workflows/build.yml/badge.svg)

The Liquibase Agent Ansible role installs [Liquibase](http://www.liquibase.org).

## :link: Requirements

See [meta/requirements.yml](meta/requirements.yml)

## :nut_and_bolt: Role Variables

* **liquibase_ver**: This property appears to specify the default version of Liquibase. **Default value** -> `4.26.0`
* **liquibase_mirror**: This property specifies the default mirror or repository from which Liquibase releases can be downloaded. **Default value** -> `https://github.com/liquibase/liquibase/releases/download`
* **liquibase_parent_install_dir**: This property indicates the default parent installation directory for Liquibase. **Default value** -> `/usr/local`
* **liquibase_checksums**: Checksums for different versions of Liquibase along with their respective download URLs. The checksums are SHA-256 hashes calculated for each Liquibase release file, ensuring the integrity of the downloaded files.

## :clipboard: Example Playbook (installs latest liquibase release)

```yml
- hosts: server
  roles:
    - role: liquibase.liquibase
```

## :clipboard: Example Playbook (installs liquibase 4.26.0)

```yml
- hosts: server
  roles:
    - role: liquibase.liquibase
      liquibase_ver:
        - 4.26.0
```
