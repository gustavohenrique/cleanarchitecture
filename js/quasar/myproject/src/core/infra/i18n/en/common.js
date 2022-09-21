export default {
  language: 'Language',
  locale: {
    en: 'English',
    pt: 'Português'
  },
  grpc: {
    error: {
      1: 'Operation was canceled',
      2: 'Our core is currently down but we are fixing it',
      3: 'Invalid argument sent',
      5: '404',
      6: 'This record already exists',
      7: 'You don\'t have enough permission',
      8: 'We have an above-normal processing that should stabilize in a few moments',
      9: 'Unable to perform this action due to some inconsistency in the data sent',
      13: 'Our core is having problems but we are already working to solve it.',
      14: 'Our database is undergoing maintenance and this may take a while',
      16: 'You must be authenticated'
    }
  },
  datetime: {
    pattern: 'yyyy-LL-dd HH:mm'
  },
  operation: {
    create: 'Create',
    list: 'List',
    back: 'Go back',
    backToView: 'Go back to view',
    backToList: 'Go back to list',
    add: 'Add',
    edit: 'Edit',
    update: 'Update',
    remove: 'Delete',
    rename: 'Rename',
    save: 'Save',
    cancel: 'Cancel',
    search: 'Search',
    view: 'View',
    loadMore: 'Load More',
    previous: 'Previous',
    next: 'Next',
    success: {
      remove: '{name} was deleted'
    },
    fail: {
      search: 'Sorry, we couldn\'t find any results',
      remove: '{name} could not be deleted'
    }
  },
  label: {
    name: 'Name',
    email: 'Email',
    about: 'About',
    createdAt: 'Created at',
    updatedAt: 'Updated at',
    revokedAt: 'Revoked at',
    expiresAt: 'Expires at',
    username: 'Username',
    sortBy: 'Sort by',
    found: 'Found',
    certificate: 'Certificate',
    error: 'Error',
    by: 'By',
    user: 'User',
    number: 'Number',
    loading: 'Loading'
  },
  notify: {
    copied: 'Copied',
    failed: 'Failed',
    noMoreResults: 'No more results'
  },
  form: {
    validation: 'Please fill correctly all required fields'
  },
  required: {
    name: 'Name is required',
    gtzero: 'Number {value} must be greater than zero'
  },
  validation: {
    min: 'Use at least {min} caracteres',
    max: 'Please use maximum {max} characters',
    required: 'Required field'
  },
  confirmation: {
    remove: 'Delete for all eternity?',
    sure: 'Are you sure?',
    removePermanently: {
      part1: 'Are you sure you want to',
      part2: 'permanently remove it?'
    },
    typeok: 'Type ok to confirm'
  },
  component: {
    searchBar: {
      searchBy: 'Search by',
      searchField: 'Search field',
      sortBy: 'Sort by'
    },
    button: {
      loadMore: 'Load more'
    },
    myPanel: {
      error: 'Something got wrong'
    }
  },
  menu: {
    projects: 'Projects',
    agents: 'Agents',
    permissions: 'Permissions',
    secrets: 'Secrets',
    support: 'Support'
  }
}
