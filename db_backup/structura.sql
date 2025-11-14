create sequence MEMBRI_SEQ
/

create sequence ABONAMENTE_SEQ
/

create sequence ANTRENORI_SEQ
/

create sequence CLASE_SEQ
/

create sequence PRODUSE_SEQ
/

create sequence ECHIPAMENTE_SEQ
/

create sequence TIPURI_ANTRENAMENT_SEQ
/

create sequence ACHIZITII_SEQ
/

create sequence ORAR_TEMPLATE_SEQ
/

create table ABONAMENTE
(
    ABONAMENTID  NUMBER(10)    not null
        primary key,
    TIPABONAMENT VARCHAR2(100) not null,
    PRET         NUMBER(10, 2) not null
)
/

create table ANTRENORI
(
    ANTRENORID   NUMBER(10)    not null
        primary key,
    NUME         VARCHAR2(100) not null,
    PRENUME      VARCHAR2(100),
    SPECIALIZARE VARCHAR2(150)
)
/

create table PRODUSE
(
    PRODUSID   NUMBER(10)    not null
        primary key,
    NUMEPRODUS VARCHAR2(150) not null,
    PRETCURENT NUMBER(10, 2),
    STOC       NUMBER(5)
)
/

create table ECHIPAMENTE
(
    ECHIPAMENTID    NUMBER(10)    not null
        primary key,
    NUMEECHIPAMENT  VARCHAR2(100) not null,
    CANTITATETOTALA NUMBER(5)
)
/

create table TIPURI_ANTRENAMENT
(
    TIPANTRENAMENTID NUMBER(10)    not null
        primary key,
    NUMEWOD          VARCHAR2(100) not null,
    DESCRIERE        VARCHAR2(1000)
)
/

create table MEMBRI
(
    MEMBRUID    NUMBER(10)    not null
        primary key,
    NUME        VARCHAR2(100) not null,
    PRENUME     VARCHAR2(100),
    EMAIL       VARCHAR2(150) not null
        unique,
    ABONAMENTID NUMBER(10)
        constraint FK_MEMBRI_ABONAMENT
            references ABONAMENTE
)
/

create table MENTORAT
(
    ANTRENORID NUMBER(10) not null
        constraint FK_MENTORAT_ANTRENOR
            references ANTRENORI,
    MEMBRUID   NUMBER(10) not null
        constraint FK_MENTORAT_MEMBRU
            references MEMBRI
                on delete cascade,
    constraint PK_MENTORAT
        primary key (ANTRENORID, MEMBRUID)
)
/

create table ACHIZITII
(
    ACHIZITIEID    NUMBER(10) not null
        primary key,
    MEMBRUID       NUMBER(10)
        constraint FK_ACHIZITII_MEMBRU
            references MEMBRI
                on delete cascade,
    PRODUSID       NUMBER(10)
        constraint FK_ACHIZITII_PRODUS
            references PRODUSE,
    DATAACHIZITIEI DATE      default SYSDATE,
    CANTITATE      NUMBER(3) default 1
)
/

create table CLASE
(
    CLASAID          NUMBER(10)    not null
        primary key,
    NUMEWOD          VARCHAR2(100) not null,
    DESCRIEREWOD     VARCHAR2(1000),
    DATAORA          TIMESTAMP(6),
    ANTRENORID       NUMBER(10)
        constraint FK_CLASE_ANTRENOR
            references ANTRENORI,
    TIPANTRENAMENTID NUMBER(10)
        constraint FK_CLASE_TIPANTRENAMENT
            references TIPURI_ANTRENAMENT
)
/

create table INSCRIERI
(
    MEMBRUID NUMBER(10) not null
        constraint FK_INSCRIERI_MEMBRU
            references MEMBRI
                on delete cascade,
    CLASAID  NUMBER(10) not null
        constraint FK_INSCRIERI_CLASA
            references CLASE
                on delete cascade,
    constraint PK_INSCRIERI
        primary key (MEMBRUID, CLASAID)
)
/

create table NECESAR_ECHIPAMENT
(
    TIPANTRENAMENTID NUMBER(10) not null
        constraint FK_NECESAR_TIP
            references TIPURI_ANTRENAMENT
                on delete cascade,
    ECHIPAMENTID     NUMBER(10) not null
        constraint FK_NECESAR_ECHIPAMENT
            references ECHIPAMENTE
                on delete cascade,
    constraint PK_NECESAR_ECHIPAMENT
        primary key (TIPANTRENAMENTID, ECHIPAMENTID)
)
/

create table ORAR_TEMPLATE
(
    TEMPLATEID       NUMBER(10)  not null
        primary key,
    ZIUASAPTAMANII   NUMBER(1)   not null
        constraint CHK_ZIUA_SAPTAMANII
            check (ZiuaSaptamanii BETWEEN 1 AND 7),
    ORA              VARCHAR2(5) not null
        constraint CHK_ORA_FORMAT
            check (REGEXP_LIKE(Ora, '^[0-2][0-9]:[0-5][0-9]$')),
    NUMEWOD_TEMPLATE VARCHAR2(100),
    ANTRENORID       NUMBER(10)
        constraint FK_ORAR_ANTRENOR
            references ANTRENORI
                on delete set null,
    TIPANTRENAMENTID NUMBER(10)
        constraint FK_ORAR_TIPANTRENAMENT
            references TIPURI_ANTRENAMENT
                on delete set null
)
/

create view V_ANTRENORI_CONTACT as
SELECT
    AntrenorID,
    Nume,
    Prenume,
    Specializare
FROM
    ANTRENORI
/

create view V_MEMBRI_ABONAMENTE as
SELECT
    m.Nume,
    m.Prenume,
    m.Email,
    a.TipAbonament,
    a.Pret
FROM
    MEMBRI m
        JOIN
    ABONAMENTE a ON m.AbonamentID = a.AbonamentID
/


