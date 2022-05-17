import $crypto from '../shared/crypto'
import $gravatar from '../shared/gravatar'
import makeUser from '../entities/userEntity'

export default class {
  toEntity (vo) {
    const entity = makeUser({
      ...vo,
      $crypto,
      $gravatar
    })
    return entity
  }

  toJSON (entity = makeUser()) {
    return {
      id: entity.getId(),
      fullName: entity.getFullName(),
      createdAt: entity.getCreatedAt(),
      email: entity.getEmail(),
      salt: entity.getSalt(),
      password: entity.getPassword(),
      picture: entity.getPicture()
    }
  }
}
