INSERT INTO Usuarios (nome, nick, email, senha)
VALUES ("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$120w6Kc8SKSuOz5g2ESV6.GzgTAOlX1H1D/biQO.Dkuv7qEpJW45W"),
       ("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$120w6Kc8SKSuOz5g2ESV6.GzgTAOlX1H1D/biQO.Dkuv7qEpJW45W"),
       ("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$120w6Kc8SKSuOz5g2ESV6.GzgTAOlX1H1D/biQO.Dkuv7qEpJW45W"),
       ("Usuario 4", "usuario_4", "usuario4@gmail.com", "$2a$10$120w6Kc8SKSuOz5g2ESV6.GzgTAOlX1H1D/biQO.Dkuv7qEpJW45W");

INSERT INTO Seguidores (usuario_id, seguidor_id)
VALUES (1, 2),
       (3, 1),
       (1, 3);

INSERT INTO Publicacoes (titulo, conteudo, autor_id)
VALUES ('Introdução ao MySQL',
        'Este artigo fornece uma introdução ao MySQL, um dos sistemas de gerenciamento de banco de dados mais populares.',
        1),
       ('Desenvolvimento de APIs com .NET', 'Aprenda como desenvolver APIs usando o framework .NET.', 2),
       ('Técnicas Avançadas de Machine Learning',
        'Este artigo explora técnicas avançadas de machine learning para projetos complexos.', 3),
       ('Boas Práticas em Engenharia de Dados',
        'Descubra as melhores práticas para trabalhar com grandes volumes de dados.', 1);

