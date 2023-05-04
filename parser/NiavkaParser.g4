parser grammar NiavkaParser;

options {
    tokenVocab=NiavkaLexer;
}

file: f_program=program EOF;

program: program_element (nl program_element)*;
program_element: diia | value | assign | nls;

diia: 'дія' d_name=identifier '(' ( nls d_params=params? nls ) ')' (d_type=type_value)? nl (d_body=body nl)? nls 'кінець';

value: NUMBER                                                   #number
     | STRING                                                   #string
     | identifier                                               #id
     | c_left=value '.' c_right=identifier                      #chain
     | c_value=value '(' c_args=args? ')'                       #call
     | '(' n_value=value ')'                                    #nested
     | a_left=value a_operation=arithmetic_op_mul a_right=value #arithmetic_mul
     | a_left=value a_operation=arithmetic_op_add a_right=value #arithmetic_add
     ;

assign: a_left=identifier (a_type=type_value)? '=' a_value=assign_value;
assign_value: value;

identifier: ID;

type_value: tv_single=identifier;

args: arg (',' arg)*;
arg: nls a_value=value nls;

params: param (nls ',' nls param)*;
param: p_name=identifier p_type=type_value;

body: body_element (nl body_element)*;
body_element: value | assign | nls;

arithmetic_op_mul: '*' | '/';
arithmetic_op_add: '+' | '-';

nl: NL;
nls: nl*;
