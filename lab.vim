inoremap ii <ESC>

" Insert mode 
inoremap <C-h> <C-o>b
inoremap <C-j> <Down>
inoremap <C-k> <Up>
inoremap <C-l> <Right>
inoremap <C-l> <C-o>w

inoremap ;ie =
inoremap ;ia +
inoremap ;il \|
inoremap ;iu _

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

vnoremap <C-v> x<Esc>"*P
nmap <Leader>pw viw"_dP