CREATE SEQUENCE COMPETITII_SEQ
/

CREATE TABLE COMPETITII
(
    COMPETITIEID NUMBER(10)    not null
        primary key,
    NUME         VARCHAR2(150) not null,
    DATA         DATE          default SYSDATE,
    LOCATIE      VARCHAR2(100),
    TAXA         NUMBER(10, 2)
)
/

CREATE TABLE PARTICIPARI_COMPETITIE
(
    COMPETITIEID NUMBER(10) not null
        constraint FK_PART_COMP
            references COMPETITII
                on delete cascade,
    MEMBRUID     NUMBER(10) not null
        constraint FK_PART_MEMBRU
            references MEMBRI
                on delete cascade,
    LOCULOBTINUT NUMBER(3),
    constraint PK_PARTICIPARI
        primary key (COMPETITIEID, MEMBRUID)
)
/