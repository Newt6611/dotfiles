vim.pack.add({
  {
    src = "https://github.com/ThePrimeagen/harpoon",
    version = "harpoon2",
  },
  {
    src = "https://github.com/nvim-lua/plenary.nvim"
  }
})

local harpoon = require("harpoon")
harpoon.setup()

local map = vim.keymap.set

map("n", "<leader>a", function() harpoon:list():add() end)
map("n", "<C-e>", function() harpoon.ui:toggle_quick_menu(harpoon:list()) end)

map("n", "<leader>y", function() harpoon:list():select(1) end)
map("n", "<leader>u", function() harpoon:list():select(2) end)
map("n", "<leader>i", function() harpoon:list():select(3) end)
map("n", "<leader>o", function() harpoon:list():select(4) end)
