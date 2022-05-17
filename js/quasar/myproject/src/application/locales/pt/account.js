const ROLE = 'Papel'
const ROLES = 'Papéis'
const GROUP = 'Grupo'
const GROUPS = `${GROUP}s`
const USER = 'Usuário'
const USERS = `${USER}s`

export default {
  plural: 'Contas',
  user: {
    singular: USER,
    plural: USERS,
    title: USERS,
    description: 'Pessoas que podem fazer o login',
    noGroup: 'Usuário não está em nenhum grupo',
    identity: 'Identidade'
  },
  role: {
    singular: ROLE,
    plural: ROLES,
    title: ROLES,
    description: 'Conjunto de responsabilidades e permissões',
    emptyPermissions: 'Você precisa adicionar pelo menos uma permissão',
    addRole: `Adicionar ${ROLE.toLowerCase()}`,
    noRoles: `Nenhum ${ROLE.toLowerCase()}`,
    addPermission: 'Adicionar permissão',
    permissions: 'Permissões',
    noPermission: 'Nenhuma permissão',
    notFound: 'A busca não retornou resultado',
    toBeAdded: 'Selecionados para serem adicionados',
    find: 'Buscar para adicionar',
    failSearch: 'A busca falhou',
    showing: 'Mostrando {total} resultado(s) na página {page}',
    target: 'Alvo',
    action: 'Ação'
  },
  group: {
    singular: GROUP,
    plural: GROUPS,
    title: GROUPS,
    description: 'Usuários herdam as permissões dos grupos do qual fazem parte',
    addPermission: 'Adicionar permissão para {name}',
    alreadyAdded: 'Permissão já foi adicionada',
    join: 'Juntar à um grupo',
    leave: 'Deixar um grupo',
    emptyRole: 'Adicione ao menos 1 papel',
    addGroup: 'Adicionar grupo',
    noRoles: 'Esse grupo não possui nenhum papél'
  },
  action: {
    read: 'Visualizar',
    create: 'Cadastrar',
    edit: 'Editar',
    remove: 'Remover',
    monitor: 'Monitorar',
    access: 'Acessar',
    block: 'Bloquear',
    run: 'Executar',
    admin: 'Administrar'
  },
  target: {
    pipeline: 'Pipeline',
    action: 'Ação',
    agent: 'Agente',
    project: 'Projeto',
    role: 'Papél',
    group: 'Grupo',
    user: 'Usuário',
    '*': 'Tudo'
  },
  label: {
    fullName: 'Nome completo',
    email: 'Email',
    add: 'Adicionar',
    remove: 'Remover'
  },
  submenu: {
    list: 'Listar',
    create: 'Cadastrar',
    edit: 'Editar',
    view: 'Visualizar',
    remove: 'Remover'
  }
}
