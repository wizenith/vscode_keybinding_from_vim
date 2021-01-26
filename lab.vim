inoremap ii <ESC>

" Insert mode 
inoremap <C-h> <C-o>b
inoremap <silent><expr><C-j> pumvisible()?"<Down>":"<c-\><c-o>gj"
inoremap <silent><expr><C-k> pumvisible()?"<Up>":"<c-\><c-o>gk"
inoremap <C-j> <Down>
inoremap <C-k> <Up>
inoremap <C-l> <Right>
inoremap <C-l> <C-o>w

inoremap <C-d>h <C-o>X
inoremap <C-d>l <C-o>x
inoremap <C-v> <C-o>"*P

inoremap ;ie =
inoremap ;ia +
inoremap ;il \|
inoremap ;iu _

inoremap ;ss $
inoremap ;sm &
inoremap ;sa @

inoremap ;it #
inoremap ;ip %
inoremap ;iq `
inoremap ;ij !
inoremap ;is *
inoremap ;ih ^
inoremap ;ik ✓
inoremap ;ii \

inoremap ;eg >=
inoremap ;el <=
inoremap ;ee ==
inoremap ;ea +=
inoremap ;en !=
inoremap ;aa ++

inoremap ;ae 1
inoremap ;ar 2
inoremap ;as 3
inoremap ;af 4
inoremap ;aw 5
inoremap ;al 6
inoremap ;ac 7
inoremap ;ab 8
inoremap ;ak 9
inoremap ;ao 0

inoremap ;n <Esc>o
inoremap ;jj <Esc>o
inoremap ;p <C-o>f)
au Filetype javascript inoremap ;j <C-o>A;<CR>
inoremap ;; <C-o>A;<CR>
inoremap ;yl <Esc>Ypi
inoremap <leader>ff <C-x><C-f>
inoremap <c-z> <C-o>u
inoremap <c-z> <Esc>ua
imap ;q <ESC>vi(S`
imap ;i' <ESC>vi(S'
imap ;i" <ESC>vi(S" 
inoremap ;ib (<Space><Space>)<C-o>h
inoremap ;iB {<Space><Space>}<C-o>h