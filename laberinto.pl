conectado(i,2).
conectado(2,3).
conectado(2,8).
conectado(8,9).
conectado(9,3).
conectado(3,4).
conectado(4,10).
conectado(10,16).
conectado(16,22).
conectado(22,21).
conectado(21,15).
conectado(15,14).
conectado(14,13).
conectado(13,7).
conectado(7,1).
conectado(14,20).
conectado(20,26).
%conectado(22,28). %Y este?
conectado(26,27).
conectado(27,28).
conectado(28,29).
conectado(29,23).
conectado(23,17).
conectado(17,11).
conectado(11,5).
conectado(5,6).
conectado(28,34).
conectado(34,35).
conectado(35,36).
conectado(36,30).
conectado(30,24).
conectado(24,18).
conectado(18,12).
conectado(34,33).
conectado(33,32).
conectado(32,31).
conectado(31,25).
conectado(25,19).
conectado(32,f).

conectados(X, Y) :- conectado(X, Y).
conectados(X, Y) :- conectado(Y, X).


ruta(Fin,Fin,LRuta,Res):-
    reverse([Fin|LRuta],Res). %% reversa al resultado obtenido
    %write("Ruta posible como solucion") ,write(Res).
    %Al final del documento se explica como ver la lista completa
    %Si no, descomente el write anterior y corrija las comas y puntos del predicado

ruta(Ini,Fin,LRuta,Res):-
    conectados(Ini,Z),
    not(member(Z,LRuta)),
    ruta(Z,Fin,[Ini|LRuta],Res).

laberinto(Ini, Fin, Ruta):-
    ruta(Ini, Fin, [], Ruta).

% Como los valores de resultado son listas con muchos valores, prolog no
% muestra todos los elementos en la respuesta, para corregir esto
% ejecute el siguiente comando antes de la llamada del predicato
% laberinto:
%     set_prolog_flag(answer_write_options, [max_depth(0)]).
