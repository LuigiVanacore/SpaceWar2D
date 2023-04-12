package airwar2d

import (
	historia "github.com/LuigiVanacore/AirWars2D/historia_engine"
	"time"
)

type Shoot struct {
	historia.SpriteNode
	duration time.Duration
	impulse  int
	rotation float64
}

func NewShoot(duration time.Duration) *Shoot {
	return &Shoot{duration: duration}
}

func (s *Shoot) Update() {
	//s.Translate()
	if s.duration <= time.Duration(0) {
		s.Delete()
	}
}

func (s *Shoot) Shoot() {

}

type PlayerShoot struct {
	Shoot
}

func NewPlayerShoot(player *Player) *PlayerShoot {
	shoot := NewShoot(time.Second * 5)
	playerShoot := &PlayerShoot{*shoot}
	position := player.GetTransform().GetPosition()
	//playerShoot.rotation = player._sprite.getRotation()/180*M_PI - M_PI/2

	playerShoot.GetTransform().SetPosition(position.X, position.Y)
	return playerShoot
}

//void Shoot::update(sf::Time deltaTime)
//{
//float seconds = deltaTime.asSeconds();
//_sprite.move(seconds * _impulse);
//_duration -= deltaTime;
//
//if(_duration < sf::Time::Zero)
//_alive = false;
//}
//
//Shoot::Shoot(Configuration::Textures textures, World& world) :Entity(textures, world)
//{
//float angle = book::random(0.f, 2.f* M_PI);
//_impulse = sf::Vector2f(std::cos(angle), std::sin(angle));
//}
///************* ShootPlayer ******************/
//
//ShootPlayer::ShootPlayer(Player& from) : Shoot(Configuration::Textures::ShootPlayer,from._world)
//{
//_duration = sf::seconds(5);
//
//float angle = from._sprite.getRotation() / 180 * M_PI - M_PI / 2;
//_impulse = sf::Vector2f(std::cos(angle),std::sin(angle)) * 500.f;
//
//setPosition(from.getPosition());
//_sprite.setRotation(from._sprite.getRotation());
//_world.add(Configuration::Sounds::LaserPlayer);
//}
//
//bool ShootPlayer::isCollide(const Entity& other)const
//{
//if(dynamic_cast<const Enemy*>(&other) != nullptr)
//{
//return Collision::circleTest(_sprite,other._sprite);
//}
//return false;
//}
//
///********************* ShootSaucer *****************/
//
//ShootSaucer::ShootSaucer(SmallSaucer& from) : Shoot(Configuration::Textures::ShootSaucer,from._world)
//{
//_duration = sf::seconds(5);
//
//
//sf::Vector2f pos = Configuration::player->getPosition() - from.getPosition();
//
//float accuracy_lost = book::random(-1.f,1.f)*M_PI/((200+Configuration::getScore())/100.f);
//float angle_rad = std::atan2(pos.y,pos.x) + accuracy_lost;
//float angle_deg = angle_rad * 180 / M_PI;
//
//_impulse = sf::Vector2f(std::cos(angle_rad),std::sin(angle_rad)) * 500.f;
//
//setPosition(from.getPosition());
//_sprite.setRotation(angle_deg + 90);
//_world.add(Configuration::Sounds::LaserEnemy);
//}
//
//bool ShootSaucer::isCollide(const Entity& other)const
//{
//if(dynamic_cast<const Player*>(&other) or dynamic_cast<const Meteor*>(&other))
//{
//return Collision::circleTest(_sprite,other._sprite);
//}
//return false;
//}
//}
