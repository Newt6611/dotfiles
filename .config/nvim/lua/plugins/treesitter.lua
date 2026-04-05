return {
  'nvim-treesitter/nvim-treesitter',
  build = ':TSUpdate',
  lazy = false,
  opts = {
    parser_install_dir = vim.fn.stdpath('data') .. '/site',
    ensure_installed = { 'lua', 'typescript', 'go', 'rust', 'solidity' },
    highlight = {
      enable = true,
    },
  },
}
