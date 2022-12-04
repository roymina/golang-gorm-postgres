 - refrence
 https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/

- to generate RSA:
  - [generate the private and public keys](http://travistidwell.com/jsencrypt/demo/)
  - [convert it to base64](https://www.base64encode.org/)
> 注意，从第一个网址生成密码对，粘贴到第二个网址编码后，填写到app.env中。共生成两个密码对。
> 从第一个网址粘贴时要全部粘贴，类似以下这种：
  ```
  -----BEGIN RSA PRIVATE KEY-----
  MIICWgIBAAKBgH3KIfK+2itJL8E9w0C42Y4AfbvV7cW9Mthmx6OsdCKhFGx/JkSY
  PtgMtEeY4J3suytFX4JOx5rjS59edrHDe5gpbaz8d5tVnuZR0F2gGWeguPCPkTsu
  G7oP002NHkHUUlUluXRkNThijhmIUNSmGRWcu6kLK+pjTlyXHuJr9/n5AgMBAAEC
  gYBTlakRQDiAYtVUttGzhCEr9q+VzQV+S8YpfcJSxBk3mYmUvriTBAdeULLKkI9Q
  4SlOC373e5mvFjH96Cs+3AKXnty4ZmfXOjtVzp+eXuJPXI670R2yIkBgrURdvNKo
  0CRyupF1qjwrlSLx88wAGcNKOOk37zhE6koq50Nl1+f0AQJBALs5jElRivD3PqLs
  pu11gpdflD9YpMccJNZxt4l2oS6HwgpRxteUMVRimatoy2BUNjiaLnuhXMiKciLq
  BAu7ZtkCQQCr/0WfUyGd8bOF2UR5ZwbJQqTDAkYf+HDchGSakS9u8tZ1LnFaxJgP
  bMHH5YtX2ufU4pK1EZ5O491fH8sZKnghAkB3WQL9hfaNDv8lurfabWs29Z26F9bK
  ej1dWhZGkZHD6JSgIWsg533erhAJfX8Pw/7gbCCvfLh5ug9yBD1aATdZAkAW2dRj
  JYVK2ajTOJlrU6/IH22KZwvwBW7hLUm8a1uU7vhlGyV+PK5DJlzcdPe9VV0FoJCD
  M/MHYiiwEaZHEqhhAkAeeQEt5cENw8sVOSFy+C5Zsget91lhjwTq3MoXnzz+wNdC
  2EJb1HjKWbWyJMWB51s0yATajfU5L2BVLdT+f9KA
  -----END RSA PRIVATE KEY-----
  ```