package database

var SEEDS = `
-- Sample data for movie database tables

-- Insert sample users
INSERT INTO users (id, username, email, password, remember_at, created_at) VALUES
(
    '550e8400-e29b-41d4-a716-446655440001',
    'moviefan123',
    'john.doe@email.com',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj/A5yPplqwu', -- password: "password123"
    NULL,
    '2024-01-15 10:30:00'
),
(
    '550e8400-e29b-41d4-a716-446655440002',
    'cinephile87',
    'sarah.smith@email.com',
    '$2a$12$XGFyP8M5Q2CjKf7Rd6YP6eOsQWXcK8NN4K9LZ3MP2Y8W1fGHx5C6a', -- password: "cinema2024"
    '2024-06-10 15:45:00',
    '2024-02-20 14:15:00'
),
(
    '550e8400-e29b-41d4-a716-446655440003',
    'filmcritic',
    'alex.review@email.com',
    '$2a$12$RKm9U7H3N2FqW6Tv8XY1BeJ4K1L5P8R7Q9S2T6V9W1C3E4F5G6H7i', -- password: "reviews456"
    NULL,
    '2024-03-05 09:20:00'
),
(
    '550e8400-e29b-41d4-a716-446655440004',
    'horrorhead',
    'mike.scary@email.com',
    '$2a$12$GH8Jp3K9L1M2N4O5P6Q7R8S9T0U1V2W3X4Y5Z6A7B8C9D0E1F2G3h', -- password: "spooky789"
    '2024-06-12 20:30:00',
    '2024-04-18 16:45:00'
);

-- Insert favorites
INSERT INTO favorites (id, user_id, imdb_id, type, created_at) VALUES
(
    '660e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440001',
    'tt0111161', -- The Shawshank Redemption
    'movie',
    '2024-01-16 11:00:00'
),
(
    '660e8400-e29b-41d4-a716-446655440002',
    '550e8400-e29b-41d4-a716-446655440001',
    'tt0468569', -- The Dark Knight
    'movie',
    '2024-01-17 13:30:00'
),
(
    '660e8400-e29b-41d4-a716-446655440003',
    '550e8400-e29b-41d4-a716-446655440001',
    'tt0944947', -- Game of Thrones
    'tv',
    '2024-01-20 19:45:00'
),
(
    '660e8400-e29b-41d4-a716-446655440004',
    '550e8400-e29b-41d4-a716-446655440002',
    'tt0110912', -- Pulp Fiction
    'movie',
    '2024-02-22 16:20:00'
),
(
    '660e8400-e29b-41d4-a716-446655440005',
    '550e8400-e29b-41d4-a716-446655440002',
    'tt0903747', -- Breaking Bad
    'tv',
    '2024-02-25 21:10:00'
),
(
    '660e8400-e29b-41d4-a716-446655440006',
    '550e8400-e29b-41d4-a716-446655440002',
    'tt1375666', -- Inception
    'movie',
    '2024-03-01 14:55:00'
),
(
    '660e8400-e29b-41d4-a716-446655440007',
    '550e8400-e29b-41d4-a716-446655440003',
    'tt0167260', -- The Lord of the Rings: The Return of the King
    'movie',
    '2024-03-07 10:15:00'
),
(
    '660e8400-e29b-41d4-a716-446655440008',
    '550e8400-e29b-41d4-a716-446655440003',
    'tt0816692', -- Interstellar
    'movie',
    '2024-03-10 18:30:00'
),
(
    '660e8400-e29b-41d4-a716-446655440009',
    '550e8400-e29b-41d4-a716-446655440004',
    'tt0108052', -- Schindler's List
    'movie',
    '2024-04-20 12:45:00'
),
(
    '660e8400-e29b-41d4-a716-446655440010',
    '550e8400-e29b-41d4-a716-446655440004',
    'tt0137523', -- Fight Club
    'movie',
    '2024-04-22 20:00:00'
);

-- Insert watchlists
INSERT INTO watchlists (id, user_id, name, created_at) VALUES
(
    '770e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440001',
    'Must Watch Movies 2024',
    '2024-01-18 12:00:00'
),
(
    '770e8400-e29b-41d4-a716-446655440002',
    '550e8400-e29b-41d4-a716-446655440001',
    'TV Series to Binge',
    '2024-02-01 15:30:00'
),
(
    '770e8400-e29b-41d4-a716-446655440003',
    '550e8400-e29b-41d4-a716-446655440002',
    'Classic Cinema',
    '2024-02-28 11:20:00'
),
(
    '770e8400-e29b-41d4-a716-446655440004',
    '550e8400-e29b-41d4-a716-446655440002',
    'Weekend Movie Night',
    '2024-03-15 19:45:00'
),
(
    '770e8400-e29b-41d4-a716-446655440005',
    '550e8400-e29b-41d4-a716-446655440003',
    'Award Winners Collection',
    '2024-03-12 14:10:00'
),
(
    '770e8400-e29b-41d4-a716-446655440006',
    '550e8400-e29b-41d4-a716-446655440004',
    'Horror Marathon',
    '2024-04-25 22:15:00'
);

-- Insert watchlist items
-- Note: Based on your schema, user_id references watchlists(id), so I'm using watchlist IDs
INSERT INTO item_watchlists (id, user_id, imdb_id, type, status, created_at) VALUES
(
    '880e8400-e29b-41d4-a716-446655440001',
    '770e8400-e29b-41d4-a716-446655440001', -- Must Watch Movies 2024
    'tt10954984', -- Ant-Man and the Wasp: Quantumania
    'movie',
    0, -- unwatched
    '2024-01-18 12:30:00'
),
(
    '880e8400-e29b-41d4-a716-446655440002',
    '770e8400-e29b-41d4-a716-446655440001',
    'tt15398776', -- Oppenheimer
    'movie',
    2, -- watched
    '2024-01-19 16:45:00'
),
(
    '880e8400-e29b-41d4-a716-446655440003',
    '770e8400-e29b-41d4-a716-446655440001',
    'tt6791350', -- Guardians of the Galaxy Vol. 3
    'movie',
    1, -- currently watching
    '2024-01-22 20:15:00'
),
(
    '880e8400-e29b-41d4-a716-446655440004',
    '770e8400-e29b-41d4-a716-446655440002', -- TV Series to Binge
    'tt2085059', -- Black Mirror
    'tv',
    0,
    '2024-02-02 10:20:00'
),
(
    '880e8400-e29b-41d4-a716-446655440005',
    '770e8400-e29b-41d4-a716-446655440002',
    'tt2861424', -- Rick and Morty
    'tv',
    1,
    '2024-02-05 18:30:00'
),
(
    '880e8400-e29b-41d4-a716-446655440006',
    '770e8400-e29b-41d4-a716-446655440002',
    'tt0141842', -- The Sopranos
    'tv',
    2,
    '2024-02-10 21:45:00'
),
(
    '880e8400-e29b-41d4-a716-446655440007',
    '770e8400-e29b-41d4-a716-446655440003', -- Classic Cinema
    'tt0050083', -- 12 Angry Men
    'movie',
    0,
    '2024-03-01 13:15:00'
),
(
    '880e8400-e29b-41d4-a716-446655440008',
    '770e8400-e29b-41d4-a716-446655440003',
    'tt0071562', -- The Godfather Part II
    'movie',
    2,
    '2024-03-05 17:20:00'
),
(
    '880e8400-e29b-41d4-a716-446655440009',
    '770e8400-e29b-41d4-a716-446655440004', -- Weekend Movie Night
    'tt0317248', -- City of God
    'movie',
    0,
    '2024-03-16 19:00:00'
),
(
    '880e8400-e29b-41d4-a716-446655440010',
    '770e8400-e29b-41d4-a716-446655440004',
    'tt0120737', -- The Lord of the Rings: The Fellowship of the Ring
    'movie',
    1,
    '2024-03-18 15:45:00'
),
(
    '880e8400-e29b-41d4-a716-446655440011',
    '770e8400-e29b-41d4-a716-446655440005', -- Award Winners Collection
    'tt1049413', -- Up
    'movie',
    2,
    '2024-03-13 11:30:00'
),
(
    '880e8400-e29b-41d4-a716-446655440012',
    '770e8400-e29b-41d4-a716-446655440005',
    'tt0993846', -- The Wolf of Wall Street
    'movie',
    0,
    '2024-03-14 16:10:00'
),
(
    '880e8400-e29b-41d4-a716-446655440013',
    '770e8400-e29b-41d4-a716-446655440006', -- Horror Marathon
    'tt0081505', -- The Shining
    'movie',
    2,
    '2024-04-26 23:00:00'
),
(
    '880e8400-e29b-41d4-a716-446655440014',
    '770e8400-e29b-41d4-a716-446655440006',
    'tt0364569', -- Oldboy
    'movie',
    1,
    '2024-04-28 22:30:00'
),
(
    '880e8400-e29b-41d4-a716-446655440015',
    '770e8400-e29b-41d4-a716-446655440006',
    'tt0078748', -- Alien
    'movie',
    0,
    '2024-05-01 21:15:00'
);
`
