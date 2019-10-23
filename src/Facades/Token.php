<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Support\Common\Facade;
use Hyperf\Extra\Contract\TokenServiceInterface;
use stdClass;

/**
 * Class Token
 * @package Hyperf\Support\Facades
 * @method static \Lcobucci\JWT\Token create(string $scene, string $jti, string $ack, array $symbol = [])
 * @method static \Lcobucci\JWT\Token get(string $tokenString)
 * @method static stdClass verify(string $scene, string $tokenString)
 */
final class Token extends Facade
{
    protected static function getFacadeAccessor()
    {
        return TokenServiceInterface::class;
    }
}