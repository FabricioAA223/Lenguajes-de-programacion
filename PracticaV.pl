% Numero 1 %
sub_cadenas(Word, Palabras, Filtradas):-
    sub_cadenas_aux(Word, Palabras, [], Filtradas).

sub_cadenas_aux(_, [], ResTemp, Res):-
    reverse(ResTemp, Res), !.

sub_cadenas_aux(Word, [W1|Ws], ResTemp, X):-
    split_string(W1, " ", "", List),
    member(Word, List),
    sub_cadenas_aux(Word, Ws, [W1|ResTemp], X), !.

sub_cadenas_aux(Word, [_|Ws], ResTemp, X):-
    sub_cadenas_aux(Word, Ws, ResTemp, X).


% Numero 2 %
sub_conjunto_aux(X,[X|_]):- !.
sub_conjunto_aux(X,[A|T]):-
    A \= X,
    sub_conjunto_aux(X,T).

sub_conjunto([],_).
sub_conjunto([X|Xs],Y):-
    sub_conjunto_aux(X,Y),
    sub_conjunto(Xs,Y).


%Numero 3 %
concatenar([], L, L).
concatenar([X|Xs], Ys, [X|Zs]) :-
    concatenar(Xs, Ys, Zs).

aplanar([],[]).
aplanar([H|T],Y):-
    H = [_|_],
    aplanar(H,Y1),
    aplanar(T,Y2),
    concatenar(Y1, Y2, Y), !.
aplanar([H|T],[H|Y]):-
    H \= [_|_],
    aplanar(T,Y).


