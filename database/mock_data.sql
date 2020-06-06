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
    'Pop is dead, long ive pop.',
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

-- save_article
INSERT INTO save_article (username, sub_name, aid) VALUES (
    'Jonny',
    'dankmeme',
    '246o1'
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
