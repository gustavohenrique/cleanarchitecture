export default function (params = {}) {
  const { $crypto, $gravatar } = params

  let id = params.id || ''
  let createdAt = params.createdAt || ''
  let fullName = params.fullName || ''
  let email = params.email || ''
  let salt = params.salt || ''
  let password = params.pasword || ''
  let picture = params.picture || ''

  function getId () {
    return id
  }

  function setId (data = '') {
    id = data
  }

  function getCreatedAt () {
    return createdAt
  }

  function setCreatedAt (data = '') {
    createdAt = data
  }

  function getFullName () {
    return fullName
  }

  function setFullName (data = '') {
    fullName = data
  }

  function getEmail () {
    return email
  }

  function setEmail (data = '') {
    email = data
  }

  function getPassword () {
    return password || ''
  }

  function setPassword (data = '') {
    password = data
  }

  function getSalt () {
    return salt
  }

  function setSalt (data = '') {
    salt = data
  }

  function getPicture () {
    if (picture) {
      return picture
    }
    return $gravatar ? $gravatar(email) : ''
  }

  function setPicture (data = '') {
    picture = data
  }

  function isValidEmail () {
    const email = getEmail()
    return email && email.indexOf('@') > 0
  }

  async function encryptPassword (rawPassword) {
    // Password deve ser um hash gerado usando pbkdf2
    // Depois criptografado com a chave publica RSA do usuario
    // Em seguida, codificado como base64
    // E no servidor decodificar e descriptografar usando a chave privada RSA
    // Salvando o hash no banco de dados
    if (!$crypto) {
      throw new Error('$crypto is not defined in userEntity')
    }
    const pbkdf2 = await $crypto.pbkdf2.hashIt({
      key: getCreatedAt(),
      salt: getSalt(),
      raw: rawPassword
    })
    const { hash, salt } = pbkdf2
    const encoded = $crypto.base64.encode(hash)
    setPassword(encoded)
    setSalt(salt)
  }
  return {
    getId,
    setId,
    getSalt,
    setSalt,
    getPassword,
    setPassword,
    getEmail,
    setEmail,
    getFullName,
    setFullName,
    getCreatedAt,
    setCreatedAt,
    isValidEmail,
    getPicture,
    setPicture,
    encryptPassword
  }
}
