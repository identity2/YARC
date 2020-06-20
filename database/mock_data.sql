-- account
INSERT INTO account (username, hashed_password, email, bio, join_time) VALUES (
    'Morrissey',
    '$2a$12$hkuxonzFkayzlECMua.l5.Cl8AHPgJ/vIBUvugFvZ7m7OXWjuT/xa',
    'morty@thesmiths.com.uk',
    'Oh mother, I can feel, the soil falling.',
    '1959-05-22 10:23:54'
);

INSERT INTO account (username, hashed_password, email, bio, join_time) VALUES (
    'Jonny',
    '$2a$12$hkuxonzFkayzlECMua.l5.Cl8AHPgJ/vIBUvugFvZ7m7OXWjuT/xa',
    'jonny@onafriday.com.uk',
    'Pop is dead, long live pop.',
    '1971-11-05 10:33:52'
);

INSERT INTO account (username, hashed_password, email, bio, join_time) VALUES (
    'Albarn',
    '$2a$12$hkuxonzFkayzlECMua.l5.Cl8AHPgJ/vIBUvugFvZ7m7OXWjuT/xa',
    'albarn@blur.com.uk',
    'Shade from the sun was his intention.',
    '1968-03-23 10:10:10'
);

-- subreddit
INSERT INTO subreddit (sub_name, description) VALUES (
    'dankmeme',
    'The dankest meme of the world!'
);

INSERT INTO subreddit (sub_name, description) VALUES (
    'meirl',
    'me irl'
);

INSERT INTO subreddit (sub_name, description) VALUES (
    'golang',
    'Hey, ho! Let''s go!'
);

INSERT INTO subreddit (sub_name, description) VALUES (
    'PHP',
    'The best programming language in the world!'
);

-- join_sub
INSERT INTO join_sub (username, sub_name) VALUES (
    'Morrissey',
    'golang'
);

INSERT INTO join_sub (username, sub_name) VALUES (
    'Jonny',
    'meirl'
);

INSERT INTO join_sub (username, sub_name) VALUES (
    'Jonny',
    'dankmeme'
);

INSERT INTO join_sub (username, sub_name) VALUES (
    'Albarn',
    'PHP'
);

-- article
INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'dankmeme',
    'WX-78',
    'text',
    'Is this a text post?',
    'An interesting title',
    'Morrissey',
    '2011-11-11 11:11:11'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'dankmeme',
    '246o1',
    'link',
    'http://www.google.com.tw/',
    'This is a nice website I discovered.',
    'Jonny',
    '2012-12-12 12:12:12'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'meirl',
    't09o3',
    'text',
    'Is this the real life, is this just fantasy? Stuck in a landslide no escape from reality. Open your eyes.',
    'meirl',
    'Albarn',
    '2016-10-30 10:30:30'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'image',
    'https://i.imgur.com/TVYXCAi.jpg',
    'This is a nice PHP tutorial I discovered',
    'Jonny',
    '2018-06-04 06:04:04'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQ',
    'text',
    'Guys, I think PHP is just too good to be real!',
    'My Confession',
    'Jonny',
    '2020-02-20 02:20:02'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'dankmeme',
    'WX-789',
    'text',
    'as title',
    'This is a text post',
    'Morrissey',
    '2011-11-18 11:11:11'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'dankmeme',
    '246o19',
    'link',
    'http://www.google.com.uk/',
    'This is a cool website I developed',
    'Jonny',
    '2012-12-13 12:12:12'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'meirl',
    't09o39',
    'text',
    'Look up to the skies and see.',
    'meirl',
    'Albarn',
    '2016-10-31 10:30:30'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'IpX1779',
    'image',
    'https://i.imgur.com/x9t4UXD.jpg',
    'This is a nice C++ tutorial I discovered',
    'Jonny',
    '2018-06-05 06:04:04'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQ9',
    'text',
    'Girls, I think Prolog is just too good to be real!',
    'One of my confessions',
    'Jonny',
    '2020-02-21 02:20:02'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQ8',
    'text',
    'This is too good to be real!',
    'confessions',
    'Jonny',
    '2020-03-21 04:20:02'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQ7',
    'text',
    'Gods, I think Rust is just too good to be real!',
    'Crabs are not smart',
    'Jonny',
    '2020-03-23 02:20:02'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQ6',
    'text',
    'Crabs, I think Perl is just too good to be real!',
    'One of my abominations',
    'Jonny',
    '2020-05-21 13:20:02'
);

INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES (
    'PHP',
    'RgMG_RTSvkQa',
    'text',
    'Abominations, I think Visual Basic is just too good to be real!',
    'One of my dominations',
    'Jonny',
    '2020-11-21 19:20:02'
);

-- comment
INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'dankmeme',
    '246o1',
    'wjifo',
    'Learn to love me, assemble the ways.',
    'Morrissey',
    '2019-09-19 09:19:19'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'dankmeme',
    '246o1',
    '01jf9',
    'I don''t know.',
    'Jonny',
    '2019-10-10 10:10:10'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'dankmeme',
    '246o1',
    'fewji3',
    'Nice.',
    'Albarn',
    '2020-01-01 01:01:01'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'meirl',
    't09o3',
    '007op',
    'Have my upvote and leave.',
    'Jonny',
    '2019-08-01 08:01:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'n00bs',
    'Yeah you are right',
    'Jonny',
    '2019-07-01 07:01:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'n01bs',
    'Epic',
    'Jonny',
    '2019-07-02 03:02:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipP',
    'That is not true',
    'Albarn',
    '2020-01-20 03:02:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipP9',
    'Wait a sec, lemme check it.',
    'Albarn',
    '2020-01-21 03:02:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipPx',
    'Yeah that is totally not true.',
    'Albarn',
    '2020-01-21 03:05:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipPz',
    'Haha, you just got roasted!',
    'Albarn',
    '2020-01-22 03:02:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipPv',
    'lmao',
    'Albarn',
    '2020-01-23 05:02:20'
);

INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES (
    'PHP',
    'IpX177',
    'WhipPc',
    'Hehe xd!',
    'Albarn',
    '2020-02-20 04:02:20'
);

-- save_article
INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Jonny',
    'dankmeme',
    '246o1'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Jonny',
    'dankmeme',
    '246o19'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Jonny',
    'meirl',
    't09o3'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Jonny',
    'PHP',
    'IpX177'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Albarn',
    'PHP',
    'IpX177'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Albarn',
    'meirl',
    't09o3'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Albarn',
    'PHP',
    'RgMG_RTSvkQ'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Albarn',
    'PHP',
    'RgMG_RTSvkQ9'
);

INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Albarn',
    'dankmeme',
    '246o19'
);

-- vote_article
INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Morrissey',
    'PHP',
    'RgMG_RTSvkQ',
    -1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Morrissey',
    'dankmeme',
    '246o1',
    -1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Jonny',
    'dankmeme',
    '246o1',
    1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Albarn',
    'dankmeme',
    '246o1',
    -1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Jonny',
    'dankmeme',
    'WX-78',
    1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Albarn',
    'dankmeme',
    'WX-78',
    1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Albarn',
    'PHP',
    'IpX177',
    1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Jonny',
    'PHP',
    'IpX177',
    1
);

INSERT INTO vote_article (username, sub_name, aid, point) VALUES (
    'Jonny',
    'PHP',
    'RgMG_RTSvkQ9',
    1
);

-- vote_comment
INSERT INTO vote_comment (username, sub_name, aid, cid, point) VALUES (
    'Jonny',
    'PHP',
    'IpX177',
    'WhipP',
    -1
);

INSERT INTO vote_comment (username, sub_name, aid, cid, point) VALUES (
    'Morrissey',
    'PHP',
    'IpX177',
    'WhipP',
    -1
);

INSERT INTO vote_comment (username, sub_name, aid, cid, point) VALUES (
    'Albarn',
    'PHP',
    'IpX177',
    'WhipP',
    -1
);

INSERT INTO vote_comment (username, sub_name, aid, cid, point) VALUES (
    'Jonny',
    'meirl',
    't09o3',
    '007op',
    1
);
