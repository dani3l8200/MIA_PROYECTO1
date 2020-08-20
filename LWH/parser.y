%{
package LWH




var root node



%}

%union{
    node node
    token string
}

/*-------------------------------TERMINALES----------------------------*/
//TOKEN PARA SEGUIR EN OTRA LINEA 
%token CONTINUE PAUSE 
//TOKENS FOR EXEC 
%token EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%token  MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%token FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 
//TOKENS FOR MOUNT 


//TOKEN PARA SEGUIR EN OTRA LINEA 
%type <token> CONTINUE PAUSE 
//TOKENS FOR EXEC 
%type <token>  EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%type <token>   MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%type <token>  FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 



//NO TERMINALES 
//OTHERS NO TERMINALS
%type <node> TYPE_NAMES
//NO TERMINALS PRINCIPALS 
%type <node> Input Command 
//NO TERMINALS FOR EXEC 
%type <node> Exec Paparams
//NO TERMINALS FOR MKDISK AND RMDIS
%type <node> Mkdisk Mkparams Rmdisk
//NO TERMINALS FOR FDISK 
%type <node> Fdisk FdiskParams


%%
//DEFENIS LA GRAMATICA
Input: /* empty */ { }
     | Command {$$ = $1 }
     
Command: EXEC Exec {$$ = Node("EXEC","exec").append($2); root = Node("EXEC","exec").append($2)}
       | MKDISK Mkdisk {$$ = Node("MKDISK",$1).append($2); root = Node("MKDISK",$1).append($2)}
       | RMDISK Rmdisk {$$ = Node("RMDISK",$1).append($2); root = Node("RMDISK", $1).append($2)}
       | FDISK Fdisk {$$ = Node("FDISK",$1).append($2); root = Node("FDISK", $1).append($2)}
       | PAUSE {$$ = Node("PAUSE",$1); root = Node("PAUSE",$1)}
       | CONTINUE {$$ = Node("COMANDO CONTINUAR",$1); root = Node("COMANDO CONTINUAR", $1)}
       ;

Exec: Exec Paparams {$$ = $1.append($2)}
    | Paparams {$$.append($1)}
    ;
Mkdisk: Mkdisk Mkparams {$$ = $1.append($2)}
      | Mkparams {$$.append($1)}
      ;
Rmdisk: Rmdisk Paparams {$$.append($1)}
      | Paparams
      ;

Fdisk: Fdisk FdiskParams {$$ = $1.append($2)}
     | FdiskParams {$$.append($1)}
     ;

FdiskParams: Paparams {$$ = $1}
           | TYPE_NAMES {$$ = $1}
           | UNIT ARROW B {$$ = Node("UNIT","B")}
           | TYPE ARROW P {$$ = Node("TYPE","P")}
           | TYPE ARROW E {$$ = Node("TYPE","E")}
           | TYPE ARROW L {$$ = Node("TYPE","L")}
           | FIT ARROW BF {$$ = Node("FIT","BF")}
           | FIT ARROW FF {$$ = Node("FIT","FF")}
           | FIT ARROW WF {$$ = Node("FIT","WF")}
           | DELETE ARROW FAST {$$ = Node("DELETE","FAST")}
           | DELETE ARROW FULL {$$ = Node("DELETE","FULL")}
           | ADD ARROW NUMBERN {$$ = Node("ADD",$3)}
           | SIZE ARROW NUMBERN {$$ = Node("SIZE",$3)}
           | UNIT ARROW K  {$$ = Node("UNIT","K")}
           | UNIT ARROW M  {$$ = Node("UNIT","M")}
           ;


Mkparams: Paparams {$$ = $1}
        | SIZE ARROW NUMBERN {$$ = Node("SIZE",$3)}
        | UNIT ARROW K  {$$ = Node("UNIT","K")}
        | UNIT ARROW M  {$$ = Node("UNIT","M")}
        | NAME ARROW ID {$$ = Node("NAME",$3)}
        ;

Paparams: PATH ARROW ROUTE {$$ = Node("PATH",$3)}
        | PATH ARROW STRTYPE {$$ = Node("PATH",$3)}
        ;

TYPE_NAMES: NAME ARROW ID {$$ = Node("NAME",$3)}
          | NAME ARROW STRTYPE {$$ = Node("NAME",$3)}





%% 
const src = `exec path`







