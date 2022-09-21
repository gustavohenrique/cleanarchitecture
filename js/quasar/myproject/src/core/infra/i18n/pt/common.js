export default {
  language: 'Idioma',
  locale: {
    en: 'English',
    pt: 'Português'
  },
  grpc: {
    error: {
      1: 'A operação foi cancelada',
      2: 'Aconteceu um problema desconhecido que já está sendo investigado',
      3: 'A informação enviada está inválida',
      5: '404',
      6: 'Já existe um cadastro com essa informação',
      7: 'Você não tem permissão suficiente',
      8: 'Estamos com um processamento acima do normal que deve estabilizar em alguns instantes',
      9: 'Não foi possível executar essa ação devido à alguma inconsistência nos dados enviados',
      13: 'Nossa aplicação está com problemas mas já estamos trabalhando para resolver isso',
      14: 'Nosso banco de dados está passando por uma manutenção e isso pode demorar um pouco',
      16: 'Você precisa estar autenticado'
    }
  },
  datetime: {
    pattern: 'dd/LL/yyyy HH:mm'
  },
  operation: {
    create: 'Cadastrar',
    list: 'Listar',
    back: 'Voltar',
    backToView: 'Voltar para visualização',
    backToList: 'Voltar para lista',
    add: 'Adicionar',
    edit: 'Editar',
    update: 'Atualizar',
    remove: 'Remover',
    rename: 'Renomear',
    save: 'Salvar',
    cancel: 'Cancelar',
    search: 'Buscar',
    view: 'Visualizar',
    loadMore: 'Carregar mais',
    previous: 'Anterior',
    next: 'Próxima',
    success: {
      remove: '{name} foi removido'
    },
    fail: {
      search: 'Desculpe, não encontramos nenhum resultado',
      remove: '{name} não pode ser removido'
    }
  },
  label: {
    name: 'Nome',
    email: 'Email',
    about: 'Sobre',
    createdAt: 'Criado em',
    updatedAt: 'Atualizado em',
    revokedAt: 'Revogado em',
    expiresAt: 'Expira em',
    sortBy: 'Ordernar por',
    found: 'Encontrado',
    certificate: 'Certificado',
    error: 'Erro',
    by: 'Por',
    user: 'Usuário',
    number: 'Número',
    loading: 'Carregando'
  },
  notify: {
    copied: 'Copiado',
    failed: 'Falhou',
    noMoreResults: 'A lista chegou ao fim'
  },
  form: {
    validation: 'Por favor, preencha corretamente todos os campos obrigatórios'
  },
  required: {
    name: 'Nome é obrigatório',
    gtzero: 'O número {value} deve ser maior que zero'
  },
  validation: {
    min: 'Preencha com pelo menos {min} caracteres',
    max: 'Não ultrapasse o máximo de {max} caracteres',
    required: 'Campo obrigatório'
  },
  confirmation: {
    remove: 'Remover por toda eternidade?',
    sure: 'Tem certeza?',
    removePermanently: {
      part1: 'Tem certeza que deseja',
      part2: 'remover permanentemente?'
    },
    typeok: 'Digite ok para confirmar'
  },
  component: {
    searchBar: {
      searchBy: 'Buscar por',
      searchField: 'Campo de busca',
      sortBy: 'Ordernar por'
    },
    button: {
      loadMore: 'Carregar mais'
    },
    myPanel: {
      error: 'Alguma coisa errada não está certa'
    }
  },
  menu: {
    projects: 'Projetos',
    agents: 'Agentes',
    permissions: 'Permissões',
    secrets: 'Segredos',
    support: 'Suporte'
  }
}
