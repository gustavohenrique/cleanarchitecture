const ROLE = 'Role'
const ROLES = `${ROLE}s`
const GROUP = 'Group'
const GROUPS = `${GROUP}s`
const USER = 'User'
const USERS = `${USER}s`

export default {
  plural: 'Accounts',
  user: {
    singular: USER,
    plural: USERS,
    title: USERS,
    description: 'People who can sign in',
    noGroup: 'This user doesn\'t have joined on any group',
    identity: 'Identity'
  },
  role: {
    singular: ROLE,
    plural: ROLES,
    title: ROLES,
    description: 'Set of permissions and rules',
    emptyPermissions: 'You must add at least one permission',
    addRole: 'Add role',
    noRoles: 'No roles',
    addPermission: 'Add permission',
    permissions: 'Permissions',
    noPermission: 'No permission',
    toBeAdded: 'Selected to be added',
    find: 'Search to add',
    failSearch: 'Search failed',
    notFound: 'No results found',
    showing: 'Showing {total} item(s) in page {page}',
    target: 'Target',
    action: 'Action'
  },
  group: {
    singular: GROUP,
    plural: GROUPS,
    title: GROUPS,
    description: 'Users inherit permissions from the groups they are joined',
    addPermission: 'Add permission to {name}',
    alreadyAdded: 'Permission already added',
    join: 'Join a group',
    leave: 'Leave group',
    emptyRole: 'Add at least 1 role',
    addGroup: 'Add group',
    noRoles: 'This group doesn\'t have any role'
  },
  action: {
    read: 'Read',
    create: 'Create',
    edit: 'Edit',
    remove: 'Remove',
    monitor: 'Monitor',
    access: 'Access',
    block: 'Block',
    run: 'Run',
    admin: 'Admin'
  },
  target: {
    pipeline: 'Pipeline',
    action: 'Action',
    agent: 'Agent',
    project: 'Project',
    role: 'Role',
    group: 'Group',
    user: 'User',
    '*': 'All'
  },
  label: {
    fullName: 'Full name',
    email: 'Email',
    add: 'Add',
    remove: 'Remove'
  },
  submenu: {
    list: 'List',
    create: 'Create',
    edit: 'Edit',
    view: 'View',
    remove: 'Remove'
  }
}
