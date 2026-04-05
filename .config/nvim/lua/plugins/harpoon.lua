return {
  {
    'ThePrimeagen/harpoon',
    branch = 'harpoon2',
    opts = {},
    keys = {
      { 
        '<leader>a', 
        mode = 'n',
        function() 
          require('harpoon'):list():add()
        end,
        desc = 'Add to harpoon',
      },
      { 
        '<C-e>', 
        mode = 'n',
        function() 
          local harpoon = require('harpoon')
          harpoon.ui:toggle_quick_menu(harpoon:list())
        end,
        desc = 'Open up harpoon',
      },
      {
        '<leader>y', 
        mode = 'n',
        function() 
          require('harpoon'):list():select(1)
        end,
      },
      {
        '<leader>u', 
        mode = 'n',
        function() 
          require('harpoon'):list():select(2)
        end,
      },
      {
        '<leader>i', 
        mode = 'n',
        function() 
          require('harpoon'):list():select(3)
        end,
      },
      {
        '<leader>o', 
        mode = 'n',
        function() 
          require('harpoon'):list():select(4)
        end,
      },
    },
  },
  { 'nvim-lua/plenary.nvim' }
}
