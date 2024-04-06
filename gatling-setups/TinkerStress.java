package tinker;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.time.Duration;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

public class TinkerStress extends Simulation {

    HttpProtocolBuilder httpProtocol = http
            .baseUrl("http://localhost:3000")
            .acceptHeader("application/json")
            .contentTypeHeader("application/json");

    // Scenario
    ScenarioBuilder scenario = scenario("Health Check")
            .exec(http("Health Check")
                    .get("/health")
                    .check(status().is(200)));

    // Configure the user load
    {
        setUp(
                scenario.injectOpen(
                        rampUsersPerSec(1).to(1000).during(Duration.ofMinutes(5)) // Aumenta gradualmente de 1 para 1000
                                                                                  // usuários por segundo durante 5
                                                                                  // minutos
                )).protocols(httpProtocol);
    }
}
